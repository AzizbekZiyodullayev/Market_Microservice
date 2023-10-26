package grpc

import (
	"genproto/product_service"

	"gitlab.com/market_mc/market_go_product_service/config"
	"gitlab.com/market_mc/market_go_product_service/grpc/client"
	"gitlab.com/market_mc/market_go_product_service/grpc/service"
	"gitlab.com/market_mc/market_go_product_service/pkg/logger"
	"gitlab.com/market_mc/market_go_product_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, srvc))
	product_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)

	return
}
