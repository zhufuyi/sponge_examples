package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"creation/internal/cache"
	"creation/internal/model"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var _ PostDao = (*postDao)(nil)

// PostDao defining the dao interface
type PostDao interface {
	Create(ctx context.Context, table *model.Post) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteByIDs(ctx context.Context, ids []uint64) error
	UpdateByID(ctx context.Context, table *model.Post) error
	GetByID(ctx context.Context, id uint64) (*model.Post, error)
	GetByCondition(ctx context.Context, condition *query.Conditions) (*model.Post, error)
	GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.Post, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Post, int64, error)

	IncrViewCount(ctx context.Context, id uint64) error
	IncrShareCount(ctx context.Context, id uint64) error

	CreateByTx(ctx context.Context, tx *gorm.DB, post *model.Post) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64, delFlag int) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Post) error

	IncrCommentCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	DecrCommentCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	IncrLikeCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	DecrLikeCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	IncrCollectCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	DecrCollectCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error
}

type postDao struct {
	db    *gorm.DB
	cache cache.PostCache
	sfg   *singleflight.Group
}

// NewPostDao creating the dao interface
func NewPostDao(db *gorm.DB, xCache cache.PostCache) PostDao {
	return &postDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

// Create a record, insert the record and the id value is written back to the table
func (d *postDao) Create(ctx context.Context, table *model.Post) error {
	err := d.db.WithContext(ctx).Create(table).Error
	_ = d.cache.Del(ctx, table.ID)
	return err
}

// DeleteByID delete a record by id
func (d *postDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Post{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, id)

	return nil
}

// DeleteByIDs delete records by batch id
func (d *postDao) DeleteByIDs(ctx context.Context, ids []uint64) error {
	err := d.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&model.Post{}).Error
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
func (d *postDao) UpdateByID(ctx context.Context, table *model.Post) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.cache.Del(ctx, table.ID)

	return err
}

func (d *postDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Post) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.PostType != 0 {
		update["post_type"] = table.PostType
	}
	if table.UserID != 0 {
		update["user_id"] = table.UserID
	}
	if table.Title != "" {
		update["title"] = table.Title
	}
	if table.Content != "" {
		update["content"] = table.Content
	}
	if table.ViewCount != 0 {
		update["view_count"] = table.ViewCount
	}
	if table.LikeCount != 0 {
		update["like_count"] = table.LikeCount
	}
	if table.CommentCount != 0 {
		update["comment_count"] = table.CommentCount
	}
	if table.CollectCount != 0 {
		update["collect_count"] = table.CollectCount
	}
	if table.ShareCount != 0 {
		update["share_count"] = table.ShareCount
	}
	if table.Longitude != 0 {
		update["longitude"] = table.Longitude
	}
	if table.Latitude != 0 {
		update["latitude"] = table.Latitude
	}
	if table.Position != "" {
		update["position"] = table.Position
	}
	if table.Visible != 0 {
		update["visible"] = table.Visible
	}
	if table.DelFlag != 0 {
		update["del_flag"] = table.DelFlag
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *postDao) GetByID(ctx context.Context, id uint64) (*model.Post, error) {
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, model.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to mysql
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.Post{}
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
			err = d.cache.Set(ctx, id, table, cache.PostExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.Post)
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
func (d *postDao) GetByCondition(ctx context.Context, c *query.Conditions) (*model.Post, error) {
	queryStr, args, err := c.ConvertToGorm()
	if err != nil {
		return nil, err
	}

	table := &model.Post{}
	err = d.db.WithContext(ctx).Where(queryStr, args...).First(table).Error
	if err != nil {
		return nil, err
	}

	return table, nil
}

// GetByIDs list of records by batch id
func (d *postDao) GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.Post, error) {
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
			var missedData []*model.Post
			err = d.db.WithContext(ctx).Where("id IN (?)", realMissedIDs).Find(&missedData).Error
			if err != nil {
				return nil, err
			}

			if len(missedData) > 0 {
				for _, data := range missedData {
					itemMap[data.ID] = data
				}
				err = d.cache.MultiSet(ctx, missedData, cache.PostExpireTime)
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
func (d *postDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Post, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.Post{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.Post{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// IncrViewCount increment view_count by 1
func (d *postDao) IncrViewCount(ctx context.Context, id uint64) error {
	err := d.db.Model(&model.Post{}).Where("id = ?", id).
		Update("view_count", gorm.Expr("view_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// IncrShareCount increment share_count by 1
func (d *postDao) IncrShareCount(ctx context.Context, id uint64) error {
	err := d.db.Model(&model.Post{}).Where("id = ?", id).
		Update("share_count", gorm.Expr("share_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// CreateByTx create a record in the database using the provided transaction
func (d *postDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Post) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record in the database using the provided transaction
func (d *postDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64, delFlag int) error {
	update := map[string]interface{}{
		"del_flag":   delFlag,
		"deleted_at": time.Now(),
	}
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).Updates(update).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.cache.Del(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *postDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Post) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.cache.Del(ctx, table.ID)

	return err
}

// IncrCommentCountByTx increment comment_count by 1 using the provided transaction
func (d *postDao) IncrCommentCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).
		Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// DecrCommentCountByTx decrement comment_count by 1 using the provided transaction
func (d *postDao) DecrCommentCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ? AND comment_count > 0", id).
		Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// IncrLikeCountByTx increment like_count by 1 using the provided transaction
func (d *postDao) IncrLikeCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// DecrLikeCountByTx decrement like_count by 1 using the provided transaction
func (d *postDao) DecrLikeCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ? AND like_count > 0", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// IncrCollectCountByTx increment collect_count by 1 using the provided transaction
func (d *postDao) IncrCollectCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).
		Update("collect_count", gorm.Expr("collect_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// DecrCollectCountByTx decrement collect_count by 1 using the provided transaction
func (d *postDao) DecrCollectCountByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Model(&model.Post{}).Where("id = ? AND collect_count > 0", id).
		Update("collect_count", gorm.Expr("collect_count - ?", 1)).Error
	if err != nil {
		return err
	}

	// delete cache
	err = d.cache.Del(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
