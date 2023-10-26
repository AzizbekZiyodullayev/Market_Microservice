package service

import (
	"context"
	"genproto/staff_service"

	"gitlab.com/market_mc/market_go_staff_service/config"
	"gitlab.com/market_mc/market_go_staff_service/grpc/client"
	"gitlab.com/market_mc/market_go_staff_service/pkg/logger"
	"gitlab.com/market_mc/market_go_staff_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StaffService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*staff_service.UnimplementedStaffServiceServer
}

func NewStaffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *StaffService {
	return &StaffService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *StaffService) Create(ctx context.Context, req *staff_service.StaffCreateReq) (*staff_service.StaffCreateResp, error) {
	u.log.Info("====== Staff Create ======", logger.Any("req", req))

	resp, err := u.strg.Staff().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating staff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (u *StaffService) GetList(ctx context.Context, req *staff_service.StaffGetListReq) (*staff_service.StaffGetListResp, error) {
	u.log.Info("====== Staff GetList ======", logger.Any("req", req))

	resp, err := u.strg.Staff().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while getting staffs", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffService) GetById(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.Staff, error) {
	u.log.Info("====== Staff GetById ======", logger.Any("req", req))

	resp, err := u.strg.Staff().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while getting staff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffService) Update(ctx context.Context, req *staff_service.StaffUpdateReq) (*staff_service.StaffUpdateResp, error) {
	u.log.Info("====== Staff Update ======", logger.Any("req", req))

	resp, err := u.strg.Staff().Update(ctx, req)
	if err != nil {
		u.log.Error("error while updating staff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffService) Delete(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.StaffDeleteResp, error) {
	u.log.Info("====== Staff Delete ======", logger.Any("req", req))

	resp, err := u.strg.Staff().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while deleting staff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
