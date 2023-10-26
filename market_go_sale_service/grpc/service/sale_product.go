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

type SaleProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*sale_service.UnimplementedSaleProductServiceServer
}

func NewSaleProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *SaleProductService {
	return &SaleProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (s *SaleProductService) Create(ctx context.Context, req *sale_service.SaleProductCreateReq) (*sale_service.SaleProductCreateResp, error) {
	s.log.Info("====== SaleProduct Create ======", logger.Any("req", req))
	resp, err := s.strg.SaleProduct().Create(ctx, req)
	if err != nil {
		s.log.Error("error while creating sale product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (s *SaleProductService) GetList(ctx context.Context, req *sale_service.SaleProductGetListReq) (*sale_service.SaleProductGetListResp, error) {
	s.log.Info("====== SaleProduct GetList ======", logger.Any("req", req))
	resp, err := s.strg.SaleProduct().GetList(ctx, req)
	if err != nil {
		s.log.Error("error while getting sale products", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (s *SaleProductService) GetById(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProduct, error) {
	s.log.Info("====== SaleProduct GetById ======", logger.Any("req", req))
	resp, err := s.strg.SaleProduct().GetById(ctx, req)
	if err != nil {
		s.log.Error("error while getting sale product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (s *SaleProductService) Update(ctx context.Context, req *sale_service.SaleProductUpdateReq) (*sale_service.SaleProductUpdateResp, error) {
	s.log.Info("====== SaleProduct Update ======", logger.Any("req", req))
	resp, err := s.strg.SaleProduct().Update(ctx, req)
	if err != nil {
		s.log.Error("error while updating sale product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (s *SaleProductService) Delete(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProductDeleteResp, error) {
	s.log.Info("====== SaleProduct Delete ======", logger.Any("req", req))
	resp, err := s.strg.SaleProduct().Delete(ctx, req)
	if err != nil {
		s.log.Error("error while deleting sale product", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
