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

type StaffTransactionService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*sale_service.UnimplementedStaffTransactionServiceServer
}

func NewStaffTransactionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *StaffTransactionService {
	return &StaffTransactionService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *StaffTransactionService) Create(ctx context.Context, req *sale_service.StaffTrCreateReq) (*sale_service.StaffTrCreateResp, error) {
	u.log.Info("====== Staff Transaction Create ======", logger.Any("req", req))

	resp, err := u.strg.StaffTransaction().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating staff transaction", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTransactionService) GetList(ctx context.Context, req *sale_service.StaffTrGetListReq) (*sale_service.StaffTrGetListResp, error) {
	u.log.Info("====== Staff Transaction GetList ======", logger.Any("req", req))

	resp, err := u.strg.StaffTransaction().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while getting staff transaction", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTransactionService) GetById(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTransaction, error) {
	u.log.Info("====== Staff Transaction GetById ======", logger.Any("req", req))

	resp, err := u.strg.StaffTransaction().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while getting staff transaction", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTransactionService) Update(ctx context.Context, req *sale_service.StaffTrUpdateReq) (*sale_service.StaffTrUpdateResp, error) {
	u.log.Info("====== Staff Transaction Update ======", logger.Any("req", req))

	resp, err := u.strg.StaffTransaction().Update(ctx, req)
	if err != nil {
		u.log.Error("error while updating staff transaction", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTransactionService) Delete(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTrDeleteResp, error) {
	u.log.Info("====== Staff Transaction Delete ======", logger.Any("req", req))

	resp, err := u.strg.StaffTransaction().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while deleting sale", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
