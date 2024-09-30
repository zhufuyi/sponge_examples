package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	"pay/internal/cache"
	"pay/internal/model"
)

var _ PayDao = (*payDao)(nil)

// PayDao defining the dao interface
type PayDao interface {
	Create(ctx context.Context, table *model.Pay) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.Pay) error
	GetByID(ctx context.Context, id uint64) (*model.Pay, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Pay, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Pay) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Pay) error

	CreatePay(tx *sql.Tx, userID uint64, orderID string, amount uint32) error
	CreatePayRevert(tx *sql.Tx, orderID string) error
}

type payDao struct {
	db    *gorm.DB
	cache cache.PayCache      // if nil, the cache is not used.
	sfg   *singleflight.Group // if cache is nil, the sfg is not used.
}

// NewPayDao creating the dao interface
func NewPayDao(db *gorm.DB, xCache cache.PayCache) PayDao {
	if xCache == nil {
		return &payDao{db: db}
	}
	return &payDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *payDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create a record, insert the record and the id value is written back to the table
func (d *payDao) Create(ctx context.Context, table *model.Pay) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID delete a record by id
func (d *payDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Pay{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID update a record by id
func (d *payDao) UpdateByID(ctx context.Context, table *model.Pay) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

func (d *payDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Pay) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.UserID != 0 {
		update["user_id"] = table.UserID
	}
	if table.OrderID != "" {
		update["order_id"] = table.OrderID
	}
	if table.Amount != 0 {
		update["amount"] = table.Amount
	}
	if table.Status != 0 {
		update["status"] = table.Status
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *payDao) GetByID(ctx context.Context, id uint64) (*model.Pay, error) {
	// no cache
	if d.cache == nil {
		record := &model.Pay{}
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
			table := &model.Pay{}
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
			err = d.cache.Set(ctx, id, table, cache.PayExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.Pay)
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
func (d *payDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Pay, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.Pay{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.Pay{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *payDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Pay) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *payDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Pay{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *payDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Pay) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

// CreatePay 创建支付订单
func (d *payDao) CreatePay(tx *sql.Tx, userId uint64, orderID string, amount uint32) error {
	sqlStr := "insert into eshop_pay.pay(user_id, order_id, amount, status, created_at, updated_at) values(?,?,?,1,now(),now())" // 1: 待支付
	sqlFieldValues := []interface{}{userId, orderID, amount}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))

	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}

// CreatePayRevert 取消支付订单
func (d *payDao) CreatePayRevert(tx *sql.Tx, orderID string) error {
	sqlStr := "update eshop_pay.pay set status=4, updated_at=now() where order_id = ?" // 4: 取消支付订单
	sqlFieldValues := []interface{}{orderID}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))

	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}