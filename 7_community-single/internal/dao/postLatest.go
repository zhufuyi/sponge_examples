package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"community/internal/cache"
	"community/internal/model"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var _ PostLatestDao = (*postLatestDao)(nil)

// PostLatestDao defining the dao interface
type PostLatestDao interface {
	Create(ctx context.Context, table *model.PostLatest) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteByIDs(ctx context.Context, ids []uint64) error
	UpdateByID(ctx context.Context, table *model.PostLatest) error
	GetByID(ctx context.Context, id uint64) (*model.PostLatest, error)
	GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.PostLatest, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.PostLatest, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, postLatest *model.PostLatest) error
	DeleteByTx(ctx context.Context, tx *gorm.DB, postID uint64, delFlag int) error
}

type postLatestDao struct {
	db    *gorm.DB
	cache cache.PostLatestCache
	sfg   *singleflight.Group
}

// NewPostLatestDao creating the dao interface
func NewPostLatestDao(db *gorm.DB, cache cache.PostLatestCache) PostLatestDao {
	return &postLatestDao{db: db, cache: cache, sfg: new(singleflight.Group)}
}

// Create a record, insert the record and the id value is written back to the table
func (d *postLatestDao) Create(ctx context.Context, table *model.PostLatest) error {
	err := d.db.WithContext(ctx).Create(table).Error
	_ = d.cache.Del(ctx, table.ID)
	return err
}

// DeleteByID delete a record based on id
func (d *postLatestDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.PostLatest{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, id)

	return nil
}

// DeleteByIDs batch delete multiple records
func (d *postLatestDao) DeleteByIDs(ctx context.Context, ids []uint64) error {
	err := d.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&model.PostLatest{}).Error
	if err != nil {
		return err
	}

	// delete cache
	for _, id := range ids {
		_ = d.cache.Del(ctx, id)
	}

	return nil
}

// UpdateByID update records by id
func (d *postLatestDao) UpdateByID(ctx context.Context, table *model.PostLatest) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.PostID != 0 {
		update["post_id"] = table.PostID
	}
	if table.UserID != 0 {
		update["user_id"] = table.UserID
	}
	if table.DelFlag != 0 {
		update["del_flag"] = table.DelFlag
	}

	err := d.db.WithContext(ctx).Model(table).Updates(update).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, table.ID)

	return nil
}

// GetByID get a record based on id
func (d *postLatestDao) GetByID(ctx context.Context, id uint64) (*model.PostLatest, error) {
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, model.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to mysql
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.PostLatest{}
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
			err = d.cache.Set(ctx, id, table, cache.PostLatestExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.PostLatest)
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

// GetByIDs get multiple rows by ids
func (d *postLatestDao) GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.PostLatest, error) {
	itemMap, err := d.cache.MultiGet(ctx, ids)
	if err != nil {
		return nil, err
	}

	var missedIDs []uint64
	for _, id := range ids {
		_, ok := itemMap[id]
		if !ok {
			missedIDs = append(missedIDs, id)
			continue
		}
	}

	// get missed data
	if len(missedIDs) > 0 {
		// find the id of an active placeholder, i.e. an id that does not exist in mysql
		var realMissedIDs []uint64
		for _, id := range missedIDs {
			_, err = d.cache.Get(ctx, id)
			if errors.Is(err, cacheBase.ErrPlaceholder) {
				continue
			} else {
				realMissedIDs = append(realMissedIDs, id)
			}
		}

		if len(realMissedIDs) > 0 {
			var missedData []*model.PostLatest
			err = d.db.WithContext(ctx).Where("id IN (?)", realMissedIDs).Find(&missedData).Error
			if err != nil {
				return nil, err
			}

			if len(missedData) > 0 {
				for _, data := range missedData {
					itemMap[data.ID] = data
				}
				err = d.cache.MultiSet(ctx, missedData, cache.PostLatestExpireTime)
				if err != nil {
					return nil, err
				}
			} else {
				for _, id := range realMissedIDs {
					_ = d.cache.SetCacheWithNotFound(ctx, id)
				}
			}
		}
	}
	return itemMap, nil
}

// GetByColumns filter multiple rows based on paging and column information
//
// params includes paging parameters and query parameters
// paging parameters (required):
//
//	page: page number, starting from 0
//	size: lines per page
//	sort: sort fields, default is id backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
//
// query parameters (not required):
//
//	name: column name
//	exp: expressions, which default to = when the value is null, have =, ! =, >, >=, <, <=, like
//	value: column name
//	logic: logical type, defaults to and when value is null, only &(and), ||(or)
//
// example: search for a male over 20 years of age
//
//	params = &query.Params{
//	    Page: 0,
//	    Size: 20,
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
func (d *postLatestDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.PostLatest, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.PostLatest{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.PostLatest{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *postLatestDao) CreateByTx(ctx context.Context, tx *gorm.DB, postLatest *model.PostLatest) error {
	return tx.WithContext(ctx).Create(postLatest).Error
}

// DeleteByTx delete a record in the database using the provided transaction
func (d *postLatestDao) DeleteByTx(ctx context.Context, tx *gorm.DB, postID uint64, delFlag int) error {
	update := map[string]interface{}{
		"del_flag":   delFlag,
		"deleted_at": time.Now(),
	}
	return tx.WithContext(ctx).Model(&model.PostLatest{}).Where("post_id = ?", postID).Updates(update).Error
}
