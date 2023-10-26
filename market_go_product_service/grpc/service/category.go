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

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*product_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *CategoryService) Create(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.CategoryCreateResp, error) {
	u.log.Info("====== Category Create ======", logger.Any("req", req))

	resp, err := u.strg.Category().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating category", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) GetList(ctx context.Context, req *product_service.CategoryGetListReq) (*product_service.CategoryGetListResp, error) {
	u.log.Info("====== Category GetList ======", logger.Any("req", req))

	resp, err := u.strg.Category().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while get list category", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) GetById(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.Category, error) {
	u.log.Info("====== Category Get ======", logger.Any("req", req))

	resp, err := u.strg.Category().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while get category", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) Update(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.CategoryUpdateResp, error) {
	u.log.Info("====== Category Update ======", logger.Any("req", req))

	resp, err := u.strg.Category().Update(ctx, req)
	if err != nil {
		u.log.Error("error while update category", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) Delete(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.CategoryDeleteResp, error) {
	u.log.Info("====== Category Delete ======", logger.Any("req", req))

	resp, err := u.strg.Category().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while delete category", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
