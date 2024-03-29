// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	community_gwV1 "community_gw/api/community_gw/v1"
	likeV1 "community_gw/api/creation/v1"
	"community_gw/internal/ecode"
	"community_gw/internal/rpcclient"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/jinzhu/copier"
)

var _ community_gwV1.LikeServiceLogicer = (*likeServiceClient)(nil)

type likeServiceClient struct {
	likeServiceCli likeV1.LikeServiceClient
}

// NewLikeServiceClient create a client
func NewLikeServiceClient() community_gwV1.LikeServiceLogicer {
	return &likeServiceClient{
		likeServiceCli: likeV1.NewLikeServiceClient(rpcclient.GetCreationRPCConn()),
	}
}

// Create 点赞
func (c *likeServiceClient) Create(ctx context.Context, req *community_gwV1.CreateLikeRequest) (*community_gwV1.CreateLikeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	_, err = c.likeServiceCli.Create(ctx, &likeV1.CreateLikeRequest{
		UserId:  req.UserId,
		ObjType: req.ObjType,
		ObjId:   req.ObjId,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.CreateLikeReply{}, nil
}

// Delete 取消点赞
func (c *likeServiceClient) Delete(ctx context.Context, req *community_gwV1.DeleteLikeRequest) (*community_gwV1.DeleteLikeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	_, err = c.likeServiceCli.Delete(ctx, &likeV1.DeleteLikeRequest{
		UserId:  req.UserId,
		ObjType: req.ObjType,
		ObjId:   req.ObjId,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.DeleteLikeReply{}, nil
}

// ListPost 获取帖子点赞列表
func (c *likeServiceClient) ListPost(ctx context.Context, req *community_gwV1.ListPostLikeRequest) (*community_gwV1.ListPostLikeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	reply, err := c.likeServiceCli.ListPost(ctx, &likeV1.ListPostLikeRequest{
		PostId: req.PostId,
		Page:   req.Page,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	likes, err := convertLikes(reply.Likes)
	if err != nil {
		return nil, ecode.StatusListPostLikeService.Err()
	}

	return &community_gwV1.ListPostLikeReply{
		Likes: likes,
		Total: reply.Total,
	}, nil
}

// ListComment 获取评论点赞列表
func (c *likeServiceClient) ListComment(ctx context.Context, req *community_gwV1.ListCommentLikeRequest) (*community_gwV1.ListCommentLikeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	reply, err := c.likeServiceCli.ListComment(ctx, &likeV1.ListCommentLikeRequest{
		CommentId: req.CommentId,
		Page:      req.Page,
		Limit:     req.Limit,
	})
	if err != nil {
		return nil, err
	}

	likes, err := convertLikes(reply.Likes)
	if err != nil {
		return nil, ecode.StatusListCommentLikeService.Err()
	}

	return &community_gwV1.ListCommentLikeReply{
		Likes: likes,
		Total: reply.Total,
	}, nil
}

func convertLike(like *likeV1.LikeInfo) (*community_gwV1.LikeInfo, error) {
	likeInfo := &community_gwV1.LikeInfo{}
	err := copier.Copy(likeInfo, like)
	return likeInfo, err
}

func convertLikes(likes []*likeV1.LikeInfo) ([]*community_gwV1.LikeInfo, error) {
	likeInfos := make([]*community_gwV1.LikeInfo, 0, len(likes))
	for _, like := range likes {
		LikeInfo, err := convertLike(like)
		if err != nil {
			return nil, err
		}
		likeInfos = append(likeInfos, LikeInfo)
	}
	return likeInfos, nil
}
