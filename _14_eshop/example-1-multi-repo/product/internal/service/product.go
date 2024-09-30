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

	productV1 "product/api/product/v1"
	"product/internal/cache"
	"product/internal/dao"
	"product/internal/ecode"
	"product/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		productV1.RegisterProductServer(server, NewProductServer()) // register service to the rpc service
	})
}

var _ productV1.ProductServer = (*product)(nil)
var _ time.Time

type product struct {
	productV1.UnimplementedProductServer

	iDao dao.ProductDao
}

// NewProductServer create a new service
func NewProductServer() productV1.ProductServer {
	return &product{
		iDao: dao.NewProductDao(
			model.GetDB(),
			cache.NewProductCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *product) Create(ctx context.Context, req *productV1.CreateProductRequest) (*productV1.CreateProductReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Product{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("product", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &productV1.CreateProductReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *product) DeleteByID(ctx context.Context, req *productV1.DeleteProductByIDRequest) (*productV1.DeleteProductByIDReply, error) {
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

	return &productV1.DeleteProductByIDReply{}, nil
}

// UpdateByID update a record by id
func (s *product) UpdateByID(ctx context.Context, req *productV1.UpdateProductByIDRequest) (*productV1.UpdateProductByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Product{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("product", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &productV1.UpdateProductByIDReply{}, nil
}

// GetByID get a record by id
func (s *product) GetByID(ctx context.Context, req *productV1.GetProductByIDRequest) (*productV1.GetProductByIDReply, error) {
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

	data, err := convertProduct(record)
	if err != nil {
		logger.Warn("convertProduct error", logger.Err(err), logger.Any("product", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDProduct.Err()
	}

	return &productV1.GetProductByIDReply{Product: data}, nil
}

// List of records by query parameters
func (s *product) List(ctx context.Context, req *productV1.ListProductRequest) (*productV1.ListProductReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	params.Size = int(req.Params.Limit)

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	products := []*productV1.Product{}
	for _, record := range records {
		data, err := convertProduct(record)
		if err != nil {
			logger.Warn("convertProduct error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		products = append(products, data)
	}

	return &productV1.ListProductReply{
		Total:    total,
		Products: products,
	}, nil
}

func convertProduct(record *model.Product) (*productV1.Product, error) {
	value := &productV1.Product{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	value.Id = record.ID
	// todo if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.CreatedAt = utils.FormatDateTimeRFC3339(record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(record.UpdatedAt)

	return value, nil
}
