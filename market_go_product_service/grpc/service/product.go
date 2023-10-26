package service

import (
	"context"
	"genproto/product_service"

	"gitlab.com/market_mc/market_go_product_service/config"
	"gitlab.com/market_mc/market_go_product_service/grpc/client"
	"gitlab.com/market_mc/market_go_product_service/pkg/logger"
	"gitlab.com/market_mc/market_go_product_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*product_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *ProductService) Create(ctx context.Context, req *product_service.ProductCreateReq) (*product_service.ProductCreateResp, error) {
	u.log.Info("====== Product Create ======", logger.Any("req", req))

	resp, err := u.strg.Product().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) GetList(ctx context.Context, req *product_service.ProductGetListReq) (*product_service.ProductGetListResp, error) {
	u.log.Info("====== Product GetList ======", logger.Any("req", req))

	resp, err := u.strg.Product().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while get list product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) GetById(ctx context.Context, req *product_service.ProductIdReq) (*product_service.Product, error) {
	u.log.Info("====== Product Get ======", logger.Any("req", req))

	resp, err := u.strg.Product().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while get product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) Update(ctx context.Context, req *product_service.ProductUpdateReq) (*product_service.ProductUpdateResp, error) {
	u.log.Info("====== Product Update ======", logger.Any("req", req))

	resp, err := u.strg.Product().Update(ctx, req)
	if err != nil {
		u.log.Error("error while update product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) Delete(ctx context.Context, req *product_service.ProductIdReq) (*product_service.ProductDeleteResp, error) {
	u.log.Info("====== Product Delete ======", logger.Any("req", req))

	resp, err := u.strg.Product().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while delete product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
