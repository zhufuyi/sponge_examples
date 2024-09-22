package handler

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	stockV1 "stock/api/stock/v1"
	"stock/internal/cache"
	"stock/internal/dao"
	"stock/internal/ecode"
	"stock/internal/model"
)

var _ stockV1.StockLogicer = (*stockHandler)(nil)
var _ time.Time

type stockHandler struct {
	stockDao dao.StockDao
}

// NewStockHandler create a handler
func NewStockHandler() stockV1.StockLogicer {
	return &stockHandler{
		stockDao: dao.NewStockDao(
			model.GetDB(),
			cache.NewStockCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (h *stockHandler) Create(ctx context.Context, req *stockV1.CreateStockRequest) (*stockV1.CreateStockReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	stock := &model.Stock{}
	err = copier.Copy(stock, req)
	if err != nil {
		return nil, ecode.ErrCreateStock.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.stockDao.Create(ctx, stock)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("stock", stock), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &stockV1.CreateStockReply{Id: stock.ID}, nil
}

// DeleteByID delete a record by id
func (h *stockHandler) DeleteByID(ctx context.Context, req *stockV1.DeleteStockByIDRequest) (*stockV1.DeleteStockByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.stockDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &stockV1.DeleteStockByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *stockHandler) UpdateByID(ctx context.Context, req *stockV1.UpdateStockByIDRequest) (*stockV1.UpdateStockByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	stock := &model.Stock{}
	err = copier.Copy(stock, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDStock.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	stock.ID = req.Id

	err = h.stockDao.UpdateByID(ctx, stock)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("stock", stock), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &stockV1.UpdateStockByIDReply{}, nil
}

// GetByID get a record by id
func (h *stockHandler) GetByID(ctx context.Context, req *stockV1.GetStockByIDRequest) (*stockV1.GetStockByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.stockDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertStock(record)
	if err != nil {
		logger.Warn("convertStock error", logger.Err(err), logger.Any("stock", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDStock.Err()
	}

	return &stockV1.GetStockByIDReply{
		Stock: data,
	}, nil
}

// List of records by query parameters
func (h *stockHandler) List(ctx context.Context, req *stockV1.ListStockRequest) (*stockV1.ListStockReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListStock.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.stockDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	stocks := []*stockV1.Stock{}
	for _, record := range records {
		data, err := convertStock(record)
		if err != nil {
			logger.Warn("convertStock error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
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
