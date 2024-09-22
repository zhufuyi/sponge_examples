package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zhufuyi/sponge/pkg/logger"
	"time"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"stock/internal/cache"
	"stock/internal/model"
)

var _ StockDao = (*stockDao)(nil)

// StockDao defining the dao interface
type StockDao interface {
	Create(ctx context.Context, table *model.Stock) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.Stock) error
	GetByID(ctx context.Context, id uint64) (*model.Stock, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Stock, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Stock) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Stock) error
}

type stockDao struct {
	db    *gorm.DB
	cache cache.StockCache    // if nil, the cache is not used.
	sfg   *singleflight.Group // if cache is nil, the sfg is not used.
}

// NewStockDao creating the dao interface
func NewStockDao(db *gorm.DB, xCache cache.StockCache) StockDao {
	if xCache == nil {
		return &stockDao{db: db}
	}
	return &stockDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *stockDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create a record, insert the record and the id value is written back to the table
func (d *stockDao) Create(ctx context.Context, table *model.Stock) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID delete a record by id
func (d *stockDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Stock{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID update a record by id
func (d *stockDao) UpdateByID(ctx context.Context, table *model.Stock) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

func (d *stockDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Stock) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.ProductID != 0 {
		update["product_id"] = table.ProductID
	}
	if table.Stock != 0 {
		update["stock"] = table.Stock
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *stockDao) GetByID(ctx context.Context, id uint64) (*model.Stock, error) {
	// no cache
	if d.cache == nil {
		record := &model.Stock{}
		err := d.db.WithContext(ctx).Where("id = ?", id).First(record).Error
		return record, err
	}

	// get from cache or database
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, model.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.Stock{}
			err = d.db.WithContext(ctx).Where("id = ?", id).First(table).Error
			if err != nil {
				// if data is empty, set not found cache to prevent cache penetration, default expiration time 10 minutes
				if errors.Is(err, model.ErrRecordNotFound) {
					err = d.cache.SetCacheWithNotFound(ctx, id)
					if err != nil {
						return nil, err
					}
					return nil, model.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			err = d.cache.Set(ctx, id, table, cache.StockExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.Stock)
		if !ok {
			return nil, model.ErrRecordNotFound
		}
		return table, nil
	} else if errors.Is(err, cacheBase.ErrPlaceholder) {
		return nil, model.ErrRecordNotFound
	}

	// fail fast, if cache error return, don't request to db
	return nil, err
}

// GetByColumns get paging records by column information,
// Note: query performance degrades when table rows are very large because of the use of offset.
//
// params includes paging parameters and query parameters
// paging parameters (required):
//
//	page: page number, starting from 0
//	limit: lines per page
//	sort: sort fields, default is id backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
//
// query parameters (not required):
//
//	name: column name
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in
//	value: column value, if exp=in, multiple values are separated by commas
//	logic: logical type, defaults to and when value is null, only &(and), ||(or)
//
// example: search for a male over 20 years of age
//
//	params = &query.Params{
//	    Page: 0,
//	    Limit: 20,
//	    Columns: []query.Column{
//		{
//			Name:    "age",
//			Exp: ">",
//			Value:   20,
//		},
//		{
//			Name:  "gender",
//			Value: "male",
//		},
//	}
func (d *stockDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Stock, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.Stock{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.Stock{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *stockDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Stock) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *stockDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Stock{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *stockDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Stock) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

// ------------------------------------------------------------------------------------------

// UpdateStockInTx update the stock of a record
func UpdateStockInTx(tx *sql.Tx, table *model.Stock) error {
	sqlStr := "update stock set stock=?, updated_at=? where id=?"
	result, err := tx.Exec(sqlStr, table.Stock, time.Now(), table.ID)
	rowCount, _ := result.RowsAffected()
	logger.Info("[mysql] info", logger.String("sql", sqlStr), logger.Any("args", []interface{}{table.Stock, time.Now(), table.ID}), logger.Int64("rows", rowCount))
	return err
}

func GetStockByID(db *sql.DB, id uint64) (string, error) {
	sqlStr := "select stock from stock where id=?"
	row := db.QueryRow(sqlStr, id)

	var stock string
	err := row.Scan(&stock)
	if err != nil {
		return "", err
	}
	logger.Info("[mysql] info", logger.String("sql", sqlStr), logger.Any("args", []interface{}{id}))
	return stock, nil
}
