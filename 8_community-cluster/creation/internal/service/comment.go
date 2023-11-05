// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"errors"

	creationV1 "creation/api/creation/v1"
	"creation/internal/cache"
	"creation/internal/dao"
	"creation/internal/ecode"
	"creation/internal/model"

	"github.com/jinzhu/copier"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		creationV1.RegisterCommentServiceServer(server, NewCommentServiceServer())
	})
}

var _ creationV1.CommentServiceServer = (*commentService)(nil)

type commentService struct {
	creationV1.UnimplementedCommentServiceServer

	postDao           dao.PostDao
	commentDao        dao.CommentDao
	commentContentDao dao.CommentContentDao
	commentHotDao     dao.CommentHotDao
	commentLatestDao  dao.CommentLatestDao
	userCommentDao    dao.UserCommentDao
}

// NewCommentServiceServer create a server
func NewCommentServiceServer() creationV1.CommentServiceServer {
	return &commentService{
		postDao: dao.NewPostDao(
			model.GetDB(),
			cache.NewPostCache(model.GetCacheType()),
		),
		commentDao: dao.NewCommentDao(
			model.GetDB(),
			cache.NewCommentCache(model.GetCacheType()),
		),
		commentContentDao: dao.NewCommentContentDao(
			model.GetDB(),
			cache.NewCommentContentCache(model.GetCacheType()),
		),
		commentHotDao: dao.NewCommentHotDao(
			model.GetDB(),
			cache.NewCommentHotCache(model.GetCacheType()),
		),
		commentLatestDao: dao.NewCommentLatestDao(
			model.GetDB(),
			cache.NewCommentLatestCache(model.GetCacheType()),
		),
		userCommentDao: dao.NewUserCommentDao(
			model.GetDB(),
			cache.NewUserCommentCache(model.GetCacheType()),
		),
	}
}

// Create 创建评论
func (s *commentService) Create(ctx context.Context, req *creationV1.CreateCommentRequest) (*creationV1.CreateCommentReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	// 判断post是否存在
	_, err = s.postDao.GetByID(ctx, req.PostId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.postDao.GetByID error", logger.Err(err), logger.Uint64("id", req.PostId), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.postDao.GetByID error", logger.Err(err), logger.Uint64("id", req.PostId), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
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

	// 创建评论
	comment := &model.Comment{
		PostID:  req.PostId,
		UserID:  req.UserId,
		DelFlag: delFlagNormal,
	}
	commentID, err := s.commentDao.CreateByTx(ctx, tx, comment)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentDao.CreateByTx error", logger.Err(err), logger.Any("comment", comment), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建评论内容
	commentContent := &model.CommentContent{
		CommentID:  commentID,
		Content:    req.Content,
		DeviceType: req.DeviceType,
		IP:         req.Ip,
	}
	_, err = s.commentContentDao.CreateByTx(ctx, tx, commentContent)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentContentDao.CreateByTx error", logger.Err(err), logger.Any("commentContent", commentContent), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建最新评论
	commentLatest := &model.CommentLatest{
		UserID:    req.UserId,
		CommentID: commentID,
		PostID:    req.PostId,
		DelFlag:   delFlagNormal,
	}
	_, err = s.commentLatestDao.CreateByTx(ctx, tx, commentLatest)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentLatestDao.CreateByTx error", logger.Err(err), logger.Any("commentLatest", commentLatest), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建热门评论
	commentHot := &model.CommentHot{
		UserID:    req.UserId,
		CommentID: commentID,
		PostID:    req.PostId,
		DelFlag:   delFlagNormal,
	}
	_, err = s.commentHotDao.CreateByTx(ctx, tx, commentHot)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentHotDao.CreateByTx error", logger.Err(err), logger.Any("commentHot", commentHot), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建用户个人评论
	userComment := &model.UserComment{
		CommentID: commentID,
		UserID:    req.UserId,
		DelFlag:   delFlagNormal,
	}
	_, err = s.userCommentDao.CreateByTx(ctx, tx, userComment)
	if err != nil {
		tx.Rollback()
		logger.Error("s.userCommentDao.CreateByTx error", logger.Err(err), logger.Any("userComment", userComment), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 评论数+1
	err = s.postDao.IncrCommentCountByTx(ctx, tx, req.PostId)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postDao.IncrCommentCountByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Error("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	commentInfo, err := convertComment(comment, commentContent)
	if err != nil {
		return nil, err
	}

	return &creationV1.CreateCommentReply{Comment: commentInfo}, nil
}

// DeleteByID 删除评论
func (s *commentService) DeleteByID(ctx context.Context, req *creationV1.DeleteCommentByIDRequest) (*creationV1.DeleteCommentByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	// 判断评论是否存在
	comment, err := s.commentDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	if comment.UserID != req.UserId {
		logger.Error("only delete your own comment", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusForbidden.Err()
	}
	if comment.DelFlag > 0 {
		return &creationV1.DeleteCommentByIDReply{}, nil
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

	// 删除评论
	err = s.commentDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除评论内容
	err = s.commentContentDao.DeleteByTx(ctx, tx, req.Id)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentContentDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除最新评论
	err = s.commentLatestDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentLatestDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除热门评论
	err = s.commentHotDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentHotDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 删除用户个人评论
	err = s.userCommentDao.DeleteByTx(ctx, tx, req.Id, int(req.DelFlag))
	if err != nil {
		tx.Rollback()
		logger.Error("s.userCommentDao.DeleteByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 评论数-1
	err = s.postDao.DecrCommentCountByTx(ctx, tx, comment.PostID)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postDao.DecrCommentCountByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Error("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.DeleteCommentByIDReply{}, nil
}

// UpdateByID 更新评论
func (s *commentService) UpdateByID(ctx context.Context, req *creationV1.UpdateCommentByIDRequest) (*creationV1.UpdateCommentByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	commentComment, err := s.commentContentDao.GetByCommentID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	commentComment.Content = req.Content
	err = s.commentContentDao.UpdateByID(ctx, commentComment)
	if err != nil {
		logger.Error("s.commentContentDao.UpdateByID error", logger.Err(err), logger.Any("commentComment", commentComment), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.UpdateCommentByIDReply{}, nil
}

// Reply 回复评论
func (s *commentService) Reply(ctx context.Context, req *creationV1.ReplyCommentRequest) (*creationV1.ReplyCommentReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	// 判断评论是否存在
	commentInfo, err := s.commentDao.GetByID(ctx, req.CommentId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("id", req.CommentId), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.commentDao.GetByID error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if commentInfo.ParentID > 0 {
		logger.Error("re-commenting on replied comments is prohibited", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusReCommentService.Err()
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

	// 创建回复评论
	comment := &model.Comment{
		PostID:   commentInfo.PostID,
		ParentID: req.CommentId,
		UserID:   req.UserId,
		DelFlag:  delFlagNormal,
	}
	commentID, err := s.commentDao.CreateByTx(ctx, tx, comment)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentDao.CreateByTx error", logger.Err(err), logger.Any("comment", comment), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建回复评论内容
	commentContent := &model.CommentContent{
		CommentID:  commentID,
		Content:    req.Content,
		DeviceType: req.DeviceType,
		IP:         req.Ip,
	}
	_, err = s.commentContentDao.CreateByTx(ctx, tx, commentContent)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentContentDao.CreateByTx error", logger.Err(err), logger.Any("commentContent", commentContent), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建最新回复评论
	commentLatest := &model.CommentLatest{
		UserID:    req.UserId,
		CommentID: commentID,
		PostID:    comment.PostID,
		ParentID:  req.CommentId,
		DelFlag:   delFlagNormal,
	}
	_, err = s.commentLatestDao.CreateByTx(ctx, tx, commentLatest)
	if err != nil {
		tx.Rollback()
		logger.Error("s.commentLatestDao.CreateByTx error", logger.Err(err), logger.Any("commentLatest", commentLatest), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 创建用户个人评论
	userComment := &model.UserComment{
		CommentID: commentID,
		UserID:    req.UserId,
		DelFlag:   delFlagNormal,
	}
	_, err = s.userCommentDao.CreateByTx(ctx, tx, userComment)
	if err != nil {
		tx.Rollback()
		logger.Error("s.userCommentDao.CreateByTx error", logger.Err(err), logger.Any("userComment", userComment), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 评论数+1
	err = s.postDao.IncrCommentCountByTx(ctx, tx, comment.PostID)
	if err != nil {
		tx.Rollback()
		logger.Error("s.postDao.IncrCommentCountByTx error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	// 评论回复数量+1
	err = s.commentDao.IncrReplyCountByTx(ctx, tx, req.CommentId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Error("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	data, err := convertComment(comment, commentContent)
	if err != nil {
		logger.Error("convertComment error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ReplyCommentReply{Comment: data}, nil
}

// GetByID 根据id获取评论
func (s *commentService) GetByID(ctx context.Context, req *creationV1.GetCommentByIDRequest) (*creationV1.GetCommentByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	comment, err := s.commentDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("commentID", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.commentDao.GetByID error", logger.Err(err), logger.Uint64("commentID", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	commentContent, err := s.commentContentDao.GetByID(ctx, comment.ID)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("s.commentContentDao.GetByID error", logger.Err(err), logger.Uint64("commentID", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.commentContentDao.GetByID error", logger.Err(err), logger.Uint64("commentID", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	data, err := convertComment(comment, commentContent)
	if err != nil {
		logger.Error("convertComment error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.GetCommentByIDReply{Comment: data}, nil
}

// ListByIDs 根据批量id获取评论列表
func (s *commentService) ListByIDs(ctx context.Context, req *creationV1.ListCommentByIDsRequest) (*creationV1.ListCommentByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	commentMap, err := s.commentDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.commentDao.GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	commentContentMap, err := s.commentContentDao.GetByCommentIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.commentContentDao.GetByCommentIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	commentInfos := []*creationV1.CommentInfo{}
	for _, id := range req.Ids {
		if comment, ok1 := commentMap[id]; ok1 {
			// 过滤已删除的评论
			if comment.DelFlag > 0 {
				continue
			}
			if commentContent, ok2 := commentContentMap[id]; ok2 {
				commentInfo, err := convertComment(comment, commentContent)
				if err != nil {
					logger.Warn("convertPost error", logger.Err(err), logger.Any("comment", comment), interceptor.ServerCtxRequestIDField(ctx))
					continue
				}
				commentInfos = append(commentInfos, commentInfo)
			}
		}
	}

	return &creationV1.ListCommentByIDsReply{Comments: commentInfos}, nil
}

// ListLatest 最新评论列表
func (s *commentService) ListLatest(ctx context.Context, req *creationV1.ListCommentLatestRequest) (*creationV1.ListCommentLatestReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.commentLatestDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-comment_id",
		Columns: []query.Column{
			{
				Name:  "post_id",
				Value: req.PostId,
			},
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.commentLatestDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListCommentLatestReply{Comments: []*creationV1.CommentInfo{}, Total: total}, nil
	}

	commentIDs := []uint64{}
	for _, record := range records {
		commentIDs = append(commentIDs, record.CommentID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListCommentByIDsRequest{Ids: commentIDs})
	if err != nil {
		logger.Warn("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListCommentLatestReply{Comments: reply.Comments, Total: total}, nil
}

// ListHot 热门评论列表
func (s *commentService) ListHot(ctx context.Context, req *creationV1.ListCommentHotRequest) (*creationV1.ListCommentHotReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.commentHotDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-comment_id",
		Columns: []query.Column{
			{
				Name:  "post_id",
				Value: req.PostId,
			},
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.commentHotDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListCommentHotReply{Comments: []*creationV1.CommentInfo{}, Total: total}, nil
	}

	commentIDs := []uint64{}
	for _, record := range records {
		commentIDs = append(commentIDs, record.CommentID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListCommentByIDsRequest{Ids: commentIDs})
	if err != nil {
		logger.Warn("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListCommentHotReply{Comments: reply.Comments, Total: total}, nil
}

// ListReply 评论回复列表
func (s *commentService) ListReply(ctx context.Context, req *creationV1.ListCommentReplyRequest) (*creationV1.ListCommentReplyReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	//ctx = interceptor.WrapServerCtx(ctx)

	records, total, err := s.commentDao.GetByColumns(ctx, &query.Params{
		Page: int(req.Page),
		Size: int(req.Limit),
		Sort: "-id",
		Columns: []query.Column{
			{
				Name:  "parent_id",
				Value: req.CommentId,
			},
			{
				Name:  "del_flag",
				Value: delFlagNormal,
			},
		},
	})
	if err != nil {
		logger.Error("s.commentDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	if len(records) == 0 {
		return &creationV1.ListCommentReplyReply{Comments: []*creationV1.CommentInfo{}, Total: total}, nil
	}

	commentIDs := []uint64{}
	for _, record := range records {
		commentIDs = append(commentIDs, record.ID)
	}

	reply, err := s.ListByIDs(ctx, &creationV1.ListCommentByIDsRequest{Ids: commentIDs})
	if err != nil {
		logger.Warn("s.ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &creationV1.ListCommentReplyReply{Comments: reply.Comments, Total: total}, nil
}

func convertComment(c *model.Comment, cc *model.CommentContent) (*creationV1.CommentInfo, error) {
	commentInfo := &creationV1.CommentInfo{}

	err := copier.Copy(commentInfo, c)
	if err != nil {
		return nil, err
	}

	// 如果字段大小不一致，需要手动转换
	commentInfo.Id = c.ID
	commentInfo.PostId = c.PostID
	commentInfo.UserId = c.UserID
	commentInfo.ParentId = c.ParentID
	commentInfo.Content = cc.Content
	commentInfo.DeviceType = cc.DeviceType
	commentInfo.Ip = cc.IP
	commentInfo.CreatedAt = c.CreatedAt.Unix()
	commentInfo.UpdatedAt = c.UpdatedAt.Unix()

	return commentInfo, nil
}
