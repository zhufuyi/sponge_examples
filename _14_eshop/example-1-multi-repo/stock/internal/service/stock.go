package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/rockscache"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
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
	db   *sql.DB
	rdb  *redis.Client
}

// NewStockServer create a new service
func NewStockServer() stockV1.StockServer {
	options := rockscache.NewDefaultOptions()
	options.StrongConsistency = true
	return &stock{
		iDao: dao.NewStockDao(
			model.GetDB(),
			cache.NewStockCache(model.GetCacheType()),
		),
		db:  model.GetSDB(),
		rdb: model.GetRedisCli(),
	}
}

// StockDeduct 扣减库存
func (s *stock) StockDeduct(ctx context.Context, req *stockV1.StockDeductRequest) (*stockV1.StockDeductReply, error) {
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
		return nil, ecode.StatusStockDeductStock.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.StockDeduct(tx, req.ProductID, req.ProductCount)
	})
	if err != nil {
		logger.Error("StockDeduct error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		if errors.Is(err, dtmcli.ErrFailure) {
			return nil, ecode.StatusAborted.ToRPCErr()
		}
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &stockV1.StockDeductReply{}, nil
}

// StockDeductRevert 补偿库存
func (s *stock) StockDeductRevert(ctx context.Context, req *stockV1.StockDeductRevertRequest) (*stockV1.StockDeductRevertReply, error) {
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
		return nil, ecode.StatusStockDeductRevertStock.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.StockDeductRevert(tx, req.ProductID, req.ProductCount)
	})
	if err != nil {
		logger.Error("CouponUseRevert error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &stockV1.StockDeductRevertReply{}, nil
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

// SetFlashSaleStock 设置秒杀产品的库存，直接更新DB和缓存，这里没有使用dtm+rockscache缓存一致性方案，主要原因是与flashSale服务使用dtm+rockscache操作redis的key相同，会产生冲突。
func (s *stock) SetFlashSaleStock(ctx context.Context, req *stockV1.SetFlashSaleStockRequest) (*stockV1.SetFlashSaleStockReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	err = s.iDao.UpdateStockByProductID(ctx, req.ProductID, req.Stock)
	if err != nil {
		logger.Warn("s.iDao.UpdateStockByProductID error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	err = s.rdb.Set(ctx, getStockCacheKey(req.ProductID), req.Stock, 24*time.Hour).Err()
	if err != nil {
		logger.Warn("s.rdb.Set error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &stockV1.SetFlashSaleStockReply{}, nil
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

func adaptErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, dtmcli.ErrFailure) {
		return ecode.StatusAborted.Err()
	}

	return ecode.StatusInternalServerError.ToRPCErr()
}

func getStockCacheKey(productID uint64) string {
	return "stock:product_id:" + strconv.FormatUint(productID, 10)
}