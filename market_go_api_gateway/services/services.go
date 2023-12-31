package services

import (
	"genproto/branch_service"
	"genproto/product_service"
	"genproto/sale_service"
	"genproto/staff_service"

	"gitlab.com/market_mc/market_go_api_gateway/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	// Product Service
	ProductService() product_service.ProductServiceClient
	CategoryService() product_service.CategoryServiceClient

	// Branch Service
	BranchService() branch_service.BranchServiceClient
	BranchProductService() branch_service.BranchProductServiceClient

	// Staff Service
	StaffTariffService() staff_service.StaffTariffServiceClient
	StaffService() staff_service.StaffServiceClient

	// Sale Service
	SaleService() sale_service.SaleServiceClient
	SaleProductService() sale_service.SaleProductServiceClient
	StaffTransactionService() sale_service.StaffTransactionServiceClient
	BranchProductTransactionService() sale_service.BrPrTransactionServiceClient
}

type grpcClients struct {
	// Product Service
	productService  product_service.ProductServiceClient
	categoryService product_service.CategoryServiceClient

	// Branch Service
	branchService        branch_service.BranchServiceClient
	branchProductService branch_service.BranchProductServiceClient

	// Staff Service
	staffTariffService staff_service.StaffTariffServiceClient
	staffService       staff_service.StaffServiceClient

	// Sale Service
	saleService                     sale_service.SaleServiceClient
	saleProductService              sale_service.SaleProductServiceClient
	staffTransactionService         sale_service.StaffTransactionServiceClient
	branchProductTransactionService sale_service.BrPrTransactionServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	// Product Microservice
	connProductService, err := grpc.Dial(
		cfg.ProductServiceHost+cfg.ProductGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// Branch Microservice
	connBranchService, err := grpc.Dial(
		cfg.BranchServiceHost+cfg.BranchGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// Staff Microservice
	connStaffService, err := grpc.Dial(
		cfg.StaffServiceHost+cfg.StaffGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// Sale Microservice
	connSaleService, err := grpc.Dial(
		cfg.SaleServiceHost+cfg.SaleGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		// Product Service
		productService:  product_service.NewProductServiceClient(connProductService),
		categoryService: product_service.NewCategoryServiceClient(connProductService),

		// Branch Service
		branchService:        branch_service.NewBranchServiceClient(connBranchService),
		branchProductService: branch_service.NewBranchProductServiceClient(connBranchService),

		// Staff Service
		staffTariffService: staff_service.NewStaffTariffServiceClient(connStaffService),
		staffService:       staff_service.NewStaffServiceClient(connStaffService),

		// Sale Service
		saleService:                     sale_service.NewSaleServiceClient(connSaleService),
		saleProductService:              sale_service.NewSaleProductServiceClient(connSaleService),
		staffTransactionService:         sale_service.NewStaffTransactionServiceClient(connSaleService),
		branchProductTransactionService: sale_service.NewBrPrTransactionServiceClient(connSaleService),
	}, nil
}

// Product Service
func (g *grpcClients) ProductService() product_service.ProductServiceClient {
	return g.productService
}

func (g *grpcClients) CategoryService() product_service.CategoryServiceClient {
	return g.categoryService
}

// Branch Service
func (g *grpcClients) BranchService() branch_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) BranchProductService() branch_service.BranchProductServiceClient {
	return g.branchProductService
}

// Staff Service
func (g *grpcClients) StaffTariffService() staff_service.StaffTariffServiceClient {
	return g.staffTariffService
}

func (g *grpcClients) StaffService() staff_service.StaffServiceClient {
	return g.staffService
}

// Sale Service
func (g *grpcClients) SaleService() sale_service.SaleServiceClient {
	return g.saleService
}

func (g *grpcClients) SaleProductService() sale_service.SaleProductServiceClient {
	return g.saleProductService
}

func (g *grpcClients) StaffTransactionService() sale_service.StaffTransactionServiceClient {
	return g.staffTransactionService
}

func (g *grpcClients) BranchProductTransactionService() sale_service.BrPrTransactionServiceClient {
	return g.branchProductTransactionService
}
