// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	creationV1 "creation/api/creation/v1"
	"creation/internal/cache"
	"creation/internal/dao"
	"creation/internal/ecode"
	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/mysql"
	"github.com/zhufuyi/sponge/pkg/mysql/query"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
)

const (
	// 帖子类型
	postTypeUnknown = 0 // 未知
	postTypeText    = 1 // 文本
	postTypeImage   = 2 // 图片
	postTypeVideo   = 3 // 视频

	// 删除帖子类型
	delFlagNormal  = 0 // 正常
	delFlagByUser  = 1 // 用户删除
	delFlagByAdmin = 2 // 管理员删除

	// 帖子可视范围
	visibleAll      = 0 // 公开
	visibleOnlySelf = 1 // 仅自己可见

	datetimeLayout = "2006-01-02 15:04:05"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		creationV1.RegisterPostServiceServer(server, NewPostServiceServer())
	})
}

var _ creationV1.PostServiceServer = (*postService)(nil)

type postService struct {
	creationV1.UnimplementedPostServiceServer

	postDao       dao.PostDao
	postLatestDao dao.PostLatestDao
	postHotDao    dao.PostHotDao
	userPostDao   dao.UserPostDao
}

// NewPostServiceServer create a server
func NewPostServiceServer() creationV1.PostServiceServer {
	return &postService{
		postDao: dao.NewPostDao(
			model.GetDB(),
			cache.NewPostCache(model.GetCacheType()),
		),
		postLatestDao: dao.NewPostLatestDao(
			model.GetDB(),
			cache.NewPostLatestCache(model.GetCacheType()),
		),
		postHotDao: dao.NewPostHotDao(
			model.GetDB(),
			cache.NewPostHotCache(model.GetCacheType()),
		),
		userPostDao: dao.NewUserPostDao(
			model.GetDB(),
			cache.NewUserPostCache(model.GetCacheType()),
		),
	}
}

// Create 创建帖子
func (s *postService) Create(ctx context.Context, req *creationV1.CreatePostRequest) (*creationV1.CreatePostReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	err = checkPostParams(req)
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	db := model.GetDB()
	tx := db.Begin()
	if tx.Error != nil {
		logger.Error("tx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	postType := getPostType(req)
	content, err := getContent(postType, req)
	if err != nil {
		logger.Warn("getContent error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, err
	}

	postData := &model.Post{
		PostType:  postType,
		UserID:    req.UserId,
		Title:     req.Title,
		Content:   content,
		Longitude: float64(req.Longitude),
		Latitude:  float64(req.Latitude),
		Position:  req.Position,
		DelFlag:   delFlagNormal,
		Visible:   visibleOnlySelf,
	}
	// 创建帖子
	postID, err := s.postDao.CreateByTx(ctx, tx, postData)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postDao.CreateByTx error", logger.Err(err), logger.Any("postData", postData), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	postLatestData := &model.PostLatest{
		PostID:  postID,
		UserID:  req.UserId,
		DelFlag: delFlagNormal,
	}
	// 创建最新帖子
	_, err = s.postLatestDao.CreateByTx(ctx, tx, postLatestData)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postLatestDao.CreateByTx error", logger.Err(err), logger.Any("postLatestData", postLatestData), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	postHotData := &model.PostHot{
		PostID:  postID,
		UserID:  req.UserId,
		DelFlag: delFlagNormal,
	}
	// 创建热门帖子
	_, err = s.postHotDao.CreateByTx(ctx, tx, postHotData)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postHotDao.CreateByTx error", logger.Err(err), logger.Any("postHotData", postHotData), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	userPostData := &model.UserPost{
		PostID:  postID,
		UserID:  req.UserId,
		DelFlag: delFlagNormal,
	}
	// 创建用户帖子
	_, err = s.userPostDao.CreateByTx(ctx, tx, userPostData)
	if err != nil {
		tx.Rollback()
		logger.Error("s.userPostDao.CreateByTx error", logger.Err(err), logger.Any("userPostData", postHotData), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Error("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	postData.ID = postID
	postInfo, err := convertPost(postData)
	if err != nil {
		// 已经添加成功，不需要返回错误
		logger.Warn("convertPost error", logger.Err(err), logger.Any("postData", postData), interceptor.ServerCtxRequestIDField(ctx))
	}

	return &creationV1.CreatePostReply{Post: postInfo}, nil
}

// UpdateContent 更新帖子内容
func (s *postService) UpdateContent(ctx context.Context, req *creationV1.UpdatePostContentRequest) (*creationV1.UpdatePostContentReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	reqVal := &creationV1.CreatePostRequest{}
	err = copier.Copy(reqVal, req)
	if err != nil {
		logger.Warn("copier.Copy error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	err = checkPostParams(reqVal)
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	postType := getPostType(reqVal)
	content, err := getContent(postType, reqVal)
	if err != nil {
		logger.Warn("getContent error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, err
	}

	err = s.postDao.UpdateByID(ctx, &model.Post{
		Model:   mysql.Model{ID: req.Id},
		Title:   req.Title,
		Content: content,
	})
	if err != nil {
		logger.Error("s.postDao.UpdateByID error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.UpdatePostContentReply{}, nil
}

// Delete 删除帖子
func (s *postService) Delete(ctx context.Context, req *creationV1.DeletePostRequest) (*creationV1.DeletePostReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	// 判断帖子是否存在
	post, err := s.postDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.postDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.postDao.GetByID error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	if post.UserID != req.UserId {
		logger.Error("only delete your own posts", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusForbidden.Err()
	}
	if post.DelFlag > 0 {
		return &creationV1.DeletePostReply{}, nil
	}

	db := model.GetDB()
	tx := db.Begin()
	if tx.Error != nil {
		logger.Error("tx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除帖子
	err = s.postDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.postDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除最新帖子
	err = s.postLatestDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.postLatestDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除热门帖子
	err = s.postHotDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.postHotDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除用户帖子
	err = s.userPostDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.userPostDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Error("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.DeletePostReply{}, nil
}

// GetByID 根据id获取帖子详情
func (s *postService) GetByID(ctx context.Context, req *creationV1.GetPostByIDRequest) (*creationV1.GetPostByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	post, err := s.postDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.postDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.postDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if post.DelFlag > 0 {
		logger.Warn("not found the post", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusNotFound.Err()
	}

	postInfo, err := convertPost(post)
	if err != nil {
		logger.Warn("convertPost error", logger.Err(err), logger.Any("post", post), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.GetPostByIDReply{Post: postInfo}, nil
}

// ListByIDs 根据批量id获取帖子列表
func (s *postService) ListByIDs(ctx context.Context, req *creationV1.ListPostByIDsRequest) (*creationV1.ListPostByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	postMap, err := s.postDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.postDao.GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	posts := []*creationV1.PostInfo{}
	for _, id := range req.Ids {
		if v, ok := postMap[id]; ok {
			// 过滤已删除的帖子
			if v.DelFlag > 0 {
				continue
			}
			post, err := convertPost(v)
			if err != nil {
				logger.Warn("convertPost error", logger.Err(err), logger.Any("post", v), interceptor.ServerCtxRequestIDField(ctx))
				continue
			}
			posts = append(posts, post)
		}
	}

	return &creationV1.ListPostByIDsReply{Posts: posts}, nil
}

// ListByUserID 我发布过的帖子列表
func (s *postService) ListByUserID(ctx context.Context, req *creationV1.ListPostByUserIDRequest) (*creationV1.ListPostByUserIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.userPostDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-post_id",
		Columns: []query.Column{
			{
				Name:  "user_id",
				Value: req.UserId,
			},
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.userPostDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListPostByUserIDReply{Posts: []*creationV1.PostInfo{}, Total: total}, nil
	}

	postIDs := []uint64{}
	for _, record := range records {
		postIDs = append(postIDs, record.PostID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListPostByIDsRequest{Ids: postIDs})
	if err != nil {
		logger.Error("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListPostByUserIDReply{Posts: reply.Posts, Total: total}, nil
}

// ListLatest 最新的帖子列表
func (s *postService) ListLatest(ctx context.Context, req *creationV1.ListPostLatestRequest) (*creationV1.ListPostLatestReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.postLatestDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-post_id",
		Columns: []query.Column{
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.postLatestDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListPostLatestReply{Posts: []*creationV1.PostInfo{}, Total: total}, nil
	}

	postIDs := []uint64{}
	for _, record := range records {
		postIDs = append(postIDs, record.PostID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListPostByIDsRequest{Ids: postIDs})
	if err != nil {
		logger.Warn("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListPostLatestReply{Posts: reply.Posts, Total: total}, nil
}

// ListHot 热门的帖子列表
func (s *postService) ListHot(ctx context.Context, req *creationV1.ListPostHotRequest) (*creationV1.ListPostHotReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.postHotDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-post_id",
		Columns: []query.Column{
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.postHotDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListPostHotReply{Posts: []*creationV1.PostInfo{}, Total: total}, nil
	}

	postIDs := []uint64{}
	for _, record := range records {
		postIDs = append(postIDs, record.PostID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListPostByIDsRequest{Ids: postIDs})
	if err != nil {
		logger.Warn("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListPostHotReply{Posts: reply.Posts, Total: total}, nil
}

// IncrViewCount 增加观看数量
func (s *postService) IncrViewCount(ctx context.Context, req *creationV1.IncrPostViewCountRequest) (*creationV1.IncrPostViewCountReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	err = s.postDao.IncrViewCount(ctx, req.Id)
	if err != nil {
		logger.Error("s.postDao.IncrViewCount error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.IncrPostViewCountReply{}, nil
}

// IncrShareCount 增加分享数量
func (s *postService) IncrShareCount(ctx context.Context, req *creationV1.IncrPostShareCountRequest) (*creationV1.IncrPostShareCountReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	err = s.postDao.IncrShareCount(ctx, req.Id)
	if err != nil {
		logger.Error("s.postDao.IncrShareCount error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.IncrPostShareCountReply{}, nil
}

func checkPostParams(req *creationV1.CreatePostRequest) error {
	if req.Text == "" && req.PicKeys == "" && req.VideoKey == "" {
		return ecode.StatusPostTypePostService.Err()
	}
	if req.PicKeys != "" && req.VideoKey != "" {
		return ecode.StatusPostType2PostService.Err()
	}
	if req.VideoKey != "" {
		if req.CoverKey == "" || req.VideoDuration == 0 || req.CoverWidth == 0 || req.CoverHeight == 0 {
			return ecode.StatusVideoParamPostService.Err()
		}
	}

	return nil
}

func getPostType(req *creationV1.CreatePostRequest) int {
	if req.PicKeys != "" {
		return postTypeImage
	}
	if req.VideoKey != "" {
		return postTypeVideo
	}
	if req.Text != "" {
		return postTypeText
	}

	return postTypeUnknown
}

func getContent(postType int, req *creationV1.CreatePostRequest) (string, error) {
	data := make(map[string]interface{})
	switch postType {
	case postTypeText:
		data["text"] = req.Text
	case postTypeImage:
		pics := strings.Split(req.PicKeys, ",")
		data["pic"] = pics
	case postTypeVideo:
		data["video"] = map[string]interface{}{
			"video_key":    req.VideoKey,
			"duration":     req.VideoDuration,
			"cover_key":    req.CoverKey,
			"cover_width":  req.CoverWidth,
			"cover_height": req.CoverHeight,
		}
	default:
		return "", ecode.StatusPostTypePostService.Err()
	}
	content, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func convertPost(post *model.Post) (*creationV1.PostInfo, error) {
	postInfo := &creationV1.PostInfo{}
	err := copier.Copy(postInfo, post)
	if err != nil {
		return nil, err
	}

	// 如果字段大小不一致，需要手动转换
	postInfo.Id = post.ID
	postInfo.UserId = post.UserID
	postInfo.CreatedAt = post.CreatedAt.Unix()
	postInfo.UpdatedAt = post.UpdatedAt.Unix()

	postInfo.Content, err = convertContent(post)
	if err != nil {
		return nil, err
	}

	return postInfo, nil
}

func convertContent(post *model.Post) (string, error) {
	if len(post.Content) == 0 {
		return "", nil
	}

	rawContent := make(map[string]interface{})
	err := json.Unmarshal([]byte(post.Content), &rawContent)
	if err != nil {
		return "", err
	}

	data := make(map[string]interface{})
	postType := post.PostType
	switch postType {
	case postTypeText:
		data["text"] = rawContent["text"]
	case postTypeImage:
		data["pic"] = rawContent["pic"]
	case postTypeVideo:
		vContent := rawContent["video"].(map[string]interface{})
		data["video"] = map[string]interface{}{
			"video_url":    vContent["video_key"],
			"duration":     vContent["duration"],
			"cover_url":    vContent["cover_key"],
			"cover_width":  vContent["cover_width"],
			"cover_height": vContent["cover_height"],
		}
	}

	content, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
