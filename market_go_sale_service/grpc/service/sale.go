package service

import (
	"context"
	"genproto/sale_service"

	"gitlab.com/market_mc/market_go_sale_service/config"
	"gitlab.com/market_mc/market_go_sale_service/grpc/client"
	"gitlab.com/market_mc/market_go_sale_service/pkg/logger"
	"gitlab.com/market_mc/market_go_sale_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SaleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*sale_service.UnimplementedSaleServiceServer
}

func NewSaleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *SaleService {
	return &SaleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *SaleService) Create(ctx context.Context, req *sale_service.SaleCreateReq) (*sale_service.SaleCreateResp, error) {
	u.log.Info("====== Sale Create ======", logger.Any("req", req))

	resp, err := u.strg.Sale().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating sale", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *SaleService) GetList(ctx context.Context, req *sale_service.SaleGetListReq) (*sale_service.SaleGetListResp, error) {
	u.log.Info("====== Sale GetList ======", logger.Any("req", req))

	resp, err := u.strg.Sale().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while getting sales", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *SaleService) GetById(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.Sale, error) {
	u.log.Info("====== Sale GetById ======", logger.Any("req", req))

	resp, err := u.strg.Sale().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while getting sale", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *SaleService) Update(ctx context.Context, req *sale_service.SaleUpdateReq) (*sale_service.SaleUpdateResp, error) {
	u.log.Info("====== Sale Update ======", logger.Any("req", req))

	resp, err := u.strg.Sale().Update(ctx, req)
	if err != nil {
		u.log.Error("error while updating sale", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *SaleService) Delete(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.SaleDeleteResp, error) {
	u.log.Info("====== Sale Delete ======", logger.Any("req", req))

	resp, err := u.strg.Sale().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while deleting sale", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
