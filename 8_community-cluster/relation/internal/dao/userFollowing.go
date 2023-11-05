package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"relation/internal/cache"
	"relation/internal/model"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var _ UserFollowingDao = (*userFollowingDao)(nil)

// UserFollowingDao defining the dao interface
type UserFollowingDao interface {
	Create(ctx context.Context, table *model.UserFollowing) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteByIDs(ctx context.Context, ids []uint64) error
	UpdateByID(ctx context.Context, table *model.UserFollowing) error
	GetByID(ctx context.Context, id uint64) (*model.UserFollowing, error)
	GetByCondition(ctx context.Context, condition *query.Conditions) (*model.UserFollowing, error)
	GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollowing, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.UserFollowing, int64, error)

	GetRelation(ctx context.Context, userID uint64, followedUid uint64) (*model.UserFollowing, error)
	BatchGetUserFollowing(ctx context.Context, userID uint64, uids []uint64) ([]*model.UserFollowing, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.UserFollowing) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.UserFollowing) error

	UpdateStatusByTx(ctx context.Context, tx *gorm.DB, userFollowing *model.UserFollowing) error
}

type userFollowingDao struct {
	db    *gorm.DB
	cache cache.UserFollowingCache
	sfg   *singleflight.Group
}

// NewUserFollowingDao creating the dao interface
func NewUserFollowingDao(db *gorm.DB, xCache cache.UserFollowingCache) UserFollowingDao {
	return &userFollowingDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

// Create a record, insert the record and the id value is written back to the table
func (d *userFollowingDao) Create(ctx context.Context, table *model.UserFollowing) error {
	err := d.db.WithContext(ctx).Create(table).Error
	_ = d.cache.Del(ctx, table.ID)
	return err
}

// DeleteByID delete a record by id
func (d *userFollowingDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.UserFollowing{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, id)

	return nil
}

// DeleteByIDs delete records by batch id
func (d *userFollowingDao) DeleteByIDs(ctx context.Context, ids []uint64) error {
	err := d.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&model.UserFollowing{}).Error
	if err != nil {
		return err
	}

	// delete cache
	for _, id := range ids {
		_ = d.cache.Del(ctx, id)
	}

	return nil
}

// UpdateByID update a record by id
func (d *userFollowingDao) UpdateByID(ctx context.Context, table *model.UserFollowing) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.cache.Del(ctx, table.ID)

	return err
}

func (d *userFollowingDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.UserFollowing) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.UserID != 0 {
		update["user_id"] = table.UserID
	}
	if table.FollowedUid != 0 {
		update["followed_uid"] = table.FollowedUid
	}
	if table.Status != 0 {
		update["status"] = table.Status
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *userFollowingDao) GetByID(ctx context.Context, id uint64) (*model.UserFollowing, error) {
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, model.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to mysql
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.UserFollowing{}
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
			err = d.cache.Set(ctx, id, table, cache.UserFollowingExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.UserFollowing)
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

// GetByCondition get a record by condition
// query conditions:
//
//	name: column name
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in
//	value: column value, if exp=in, multiple values are separated by commas
//	logic: logical type, defaults to and when value is null, only &(and), ||(or)
//
// example: find a male aged 20
//
//	condition = &query.Conditions{
//	    Columns: []query.Column{
//		{
//			Name:    "age",
//			Value:   20,
//		},
//		{
//			Name:  "gender",
//			Value: "male",
//		},
//	}
func (d *userFollowingDao) GetByCondition(ctx context.Context, c *query.Conditions) (*model.UserFollowing, error) {
	queryStr, args, err := c.ConvertToGorm()
	if err != nil {
		return nil, err
	}

	table := &model.UserFollowing{}
	err = d.db.WithContext(ctx).Where(queryStr, args...).First(table).Error
	if err != nil {
		return nil, err
	}

	return table, nil
}

// GetByIDs list of records by batch id
func (d *userFollowingDao) GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollowing, error) {
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
			}
			realMissedIDs = append(realMissedIDs, id)
		}

		if len(realMissedIDs) > 0 {
			var missedData []*model.UserFollowing
			err = d.db.WithContext(ctx).Where("id IN (?)", realMissedIDs).Find(&missedData).Error
			if err != nil {
				return nil, err
			}

			if len(missedData) > 0 {
				for _, data := range missedData {
					itemMap[data.ID] = data
				}
				err = d.cache.MultiSet(ctx, missedData, cache.UserFollowingExpireTime)
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

// GetByColumns get records by paging and column information,
// Note: query performance degrades when table rows are very large because of the use of offset.
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
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in
//	value: column value, if exp=in, multiple values are separated by commas
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
func (d *userFollowingDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.UserFollowing, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.UserFollowing{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.UserFollowing{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *userFollowingDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.UserFollowing) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *userFollowingDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	update := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	err := tx.WithContext(ctx).Model(&model.UserFollowing{}).Where("id = ?", id).Updates(update).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *userFollowingDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.UserFollowing) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.cache.Del(ctx, table.ID)

	return err
}

// GetRelation get user relation
func (d *userFollowingDao) GetRelation(ctx context.Context, userID uint64, followedUid uint64) (*model.UserFollowing, error) {
	userFollowing := &model.UserFollowing{}
	err := d.db.WithContext(ctx).Where("user_id = ? AND followed_uid = ?", userID, followedUid).First(userFollowing).Error
	return userFollowing, err
}

// UpdateStatusByTx unfollow user using the provided transaction
func (d *userFollowingDao) UpdateStatusByTx(ctx context.Context, tx *gorm.DB, userFollowing *model.UserFollowing) error {
	update := map[string]interface{}{
		"status": userFollowing.Status,
	}

	return tx.WithContext(ctx).Model(userFollowing).Where("user_id = ? AND followed_uid = ?", userFollowing.UserID, userFollowing.FollowedUid).Updates(update).Error
}

func (d *userFollowingDao) BatchGetUserFollowing(ctx context.Context, userID uint64, uids []uint64) ([]*model.UserFollowing, error) {
	userFollowings := make([]*model.UserFollowing, 0)
	err := d.db.WithContext(ctx).Where("user_id=? AND followed_uid in (?) and status=1", userID, uids).Find(&userFollowings).Error
	if err != nil {
		return nil, err
	}

	return userFollowings, nil
}
