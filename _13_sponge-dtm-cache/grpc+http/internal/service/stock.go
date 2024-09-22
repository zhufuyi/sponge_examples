package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	stockV1 "stock/api/stock/v1"
	"stock/internal/cache"
	"stock/internal/dao"
	"stock/internal/ecode"
	"stock/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		stockV1.RegisterStockServer(server, NewStockServer()) // register service to the rpc service
	})
}

var _ stockV1.StockServer = (*stock)(nil)
var _ time.Time

type stock struct {
	stockV1.UnimplementedStockServer

	iDao dao.StockDao
}

// NewStockServer create a new service
func NewStockServer() stockV1.StockServer {
	return &stock{
		iDao: dao.NewStockDao(
			model.GetDB(),
			cache.NewStockCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *stock) Create(ctx context.Context, req *stockV1.CreateStockRequest) (*stockV1.CreateStockReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Stock{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateStock.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("stock", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &stockV1.CreateStockReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *stock) DeleteByID(ctx context.Context, req *stockV1.DeleteStockByIDRequest) (*stockV1.DeleteStockByIDReply, error) {
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

	return &stockV1.DeleteStockByIDReply{}, nil
}

// UpdateByID update a record by id
func (s *stock) UpdateByID(ctx context.Context, req *stockV1.UpdateStockByIDRequest) (*stockV1.UpdateStockByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Stock{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDStock.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("stock", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &stockV1.UpdateStockByIDReply{}, nil
}

// GetByID get a record by id
func (s *stock) GetByID(ctx context.Context, req *stockV1.GetStockByIDRequest) (*stockV1.GetStockByIDReply, error) {
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

	data, err := convertStock(record)
	if err != nil {
		logger.Warn("convertStock error", logger.Err(err), logger.Any("stock", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDStock.Err()
	}

	return &stockV1.GetStockByIDReply{Stock: data}, nil
}

// List of records by query parameters
func (s *stock) List(ctx context.Context, req *stockV1.ListStockRequest) (*stockV1.ListStockReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListStock.Err()
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

	stocks := []*stockV1.Stock{}
	for _, record := range records {
		data, err := convertStock(record)
		if err != nil {
			logger.Warn("convertStock error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		stocks = append(stocks, data)
	}

	return &stockV1.ListStockReply{
		Total:  total,
		Stocks: stocks,
	}, nil
}

func convertStock(record *model.Stock) (*stockV1.Stock, error) {
	value := &stockV1.Stock{}
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
