// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	community_gwV1 "community_gw/api/community_gw/v1"
	relationV1 "community_gw/api/relation/v1"
	"community_gw/internal/ecode"
	"community_gw/internal/rpcclient"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
)

var _ community_gwV1.RelationServiceLogicer = (*relationServiceClient)(nil)

type relationServiceClient struct {
	relationServiceCli relationV1.RelationServiceClient
}

// NewRelationServiceClient create a client
func NewRelationServiceClient() community_gwV1.RelationServiceLogicer {
	return &relationServiceClient{
		relationServiceCli: relationV1.NewRelationServiceClient(rpcclient.GetRelationRPCConn()),
	}
}

// Follow 关注
func (c *relationServiceClient) Follow(ctx context.Context, req *community_gwV1.FollowRequest) (*community_gwV1.FollowReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	_, err = c.relationServiceCli.Follow(ctx, &relationV1.FollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.FollowReply{}, nil
}

// Unfollow 取消关注
func (c *relationServiceClient) Unfollow(ctx context.Context, req *community_gwV1.UnfollowRequest) (*community_gwV1.UnfollowReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	_, err = c.relationServiceCli.Unfollow(ctx, &relationV1.UnfollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.UnfollowReply{}, nil
}

// ListFollowing 关注列表
func (c *relationServiceClient) ListFollowing(ctx context.Context, req *community_gwV1.ListFollowingRequest) (*community_gwV1.ListFollowingReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	reply, err := c.relationServiceCli.ListFollowing(ctx, &relationV1.ListFollowingRequest{
		UserId: req.UserId,
		Page:   req.Page,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.ListFollowingReply{
		FollowedUids: reply.FollowedUids,
		Total:        reply.Total,
	}, nil
}

// ListFollower 粉丝列表
func (c *relationServiceClient) ListFollower(ctx context.Context, req *community_gwV1.ListFollowerRequest) (*community_gwV1.ListFollowerReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	reply, err := c.relationServiceCli.ListFollower(ctx, &relationV1.ListFollowerRequest{
		UserId: req.UserId,
		Page:   req.Page,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.ListFollowerReply{
		FollowerUids: reply.FollowerUids,
		Total:        reply.Total,
	}, nil
}

// BatchGetRelation 批量获取关注关系，a和b,c,d的关注状态
func (c *relationServiceClient) BatchGetRelation(ctx context.Context, req *community_gwV1.BatchGetRelationRequest) (*community_gwV1.BatchGetRelationReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	reply, err := c.relationServiceCli.BatchGetRelation(ctx, &relationV1.BatchGetRelationRequest{
		UserId: req.UserId,
		Uids:   req.Uids,
	})
	if err != nil {
		return nil, err
	}

	return &community_gwV1.BatchGetRelationReply{
		Result: reply.Result,
	}, nil
}