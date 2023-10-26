package grpc

import (
	"genproto/branch_service"

	"gitlab.com/market_mc/market_go_branch_service/config"
	"gitlab.com/market_mc/market_go_branch_service/grpc/client"
	"gitlab.com/market_mc/market_go_branch_service/grpc/service"
	"gitlab.com/market_mc/market_go_branch_service/pkg/logger"
	"gitlab.com/market_mc/market_go_branch_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	branch_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)

	return
}
