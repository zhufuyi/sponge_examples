package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	couponV1 "coupon/api/coupon/v1"
	"coupon/internal/cache"
	"coupon/internal/dao"
	"coupon/internal/ecode"
	"coupon/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		couponV1.RegisterCouponServer(server, NewCouponServer()) // register service to the rpc service
	})
}

var _ couponV1.CouponServer = (*coupon)(nil)
var _ time.Time

type coupon struct {
	couponV1.UnimplementedCouponServer

	iDao dao.CouponDao
}

// NewCouponServer create a new service
func NewCouponServer() couponV1.CouponServer {
	return &coupon{
		iDao: dao.NewCouponDao(
			model.GetDB(),
			cache.NewCouponCache(model.GetCacheType()),
		),
	}
}

// CouponUse 使用优惠券
func (s *coupon) CouponUse(ctx context.Context, req *couponV1.CouponUseRequest) (*couponV1.CouponUseReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	// 使用子事务屏障
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		logger.Error("BarrierFromGrpc error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusCouponUseCoupon.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CouponUse(tx, req.CouponID)
	})
	if err != nil {
		logger.Error("CouponUse error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		if errors.Is(err, dtmcli.ErrFailure) {
			return nil, ecode.StatusAborted.ToRPCErr(dtmcli.ResultFailure)
		}
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &couponV1.CouponUseReply{}, nil
}

// CouponUseRevert 补偿优惠券
func (s *coupon) CouponUseRevert(ctx context.Context, req *couponV1.CouponUseRevertRequest) (*couponV1.CouponUseRevertReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	// 使用子事务屏障
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		logger.Error("BarrierFromGrpc error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusCouponUseCoupon.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CouponUseRevert(tx, req.CouponID)
	})
	if err != nil {
		logger.Error("CouponUseRevert error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &couponV1.CouponUseRevertReply{}, nil
}

// Create a record
func (s *coupon) Create(ctx context.Context, req *couponV1.CreateCouponRequest) (*couponV1.CreateCouponReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Coupon{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("coupon", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &couponV1.CreateCouponReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *coupon) DeleteByID(ctx context.Context, req *couponV1.DeleteCouponByIDRequest) (*couponV1.DeleteCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &couponV1.DeleteCouponByIDReply{}, nil
}

// UpdateByID update a record by id
func (s *coupon) UpdateByID(ctx context.Context, req *couponV1.UpdateCouponByIDRequest) (*couponV1.UpdateCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Coupon{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("coupon", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &couponV1.UpdateCouponByIDReply{}, nil
}

// GetByID get a record by id
func (s *coupon) GetByID(ctx context.Context, req *couponV1.GetCouponByIDRequest) (*couponV1.GetCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertCoupon(record)
	if err != nil {
		logger.Warn("convertCoupon error", logger.Err(err), logger.Any("coupon", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDCoupon.Err()
	}

	return &couponV1.GetCouponByIDReply{Coupon: data}, nil
}

// List of records by query parameters
func (s *coupon) List(ctx context.Context, req *couponV1.ListCouponRequest) (*couponV1.ListCouponReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	coupons := []*couponV1.Coupon{}
	for _, record := range records {
		data, err := convertCoupon(record)
		if err != nil {
			logger.Warn("convertCoupon error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		coupons = append(coupons, data)
	}

	return &couponV1.ListCouponReply{
		Total:   total,
		Coupons: coupons,
	}, nil
}

func convertCoupon(record *model.Coupon) (*couponV1.Coupon, error) {
	value := &couponV1.Coupon{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(*record.UpdatedAt)

	return value, nil
}
