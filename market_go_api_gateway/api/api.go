package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/market_mc/market_go_api_gateway/api/docs"
	"gitlab.com/market_mc/market_go_api_gateway/api/handler"
	"gitlab.com/market_mc/market_go_api_gateway/config"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func SetUpApi(r *gin.Engine, h *handler.Handler, cfg config.Config) {
	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(500))

	v1 := r.Group("/v1")

	// branch api
	v1.POST("/branches", h.CreateBranch)
	v1.GET("/branches", h.GetListBranch)
	v1.GET("/branches/:id", h.GetBranch)
	v1.PUT("/branches/:id", h.UpdateBranch)
	v1.DELETE("/branches/:id", h.DeleteBranch)

	// branch product api
	v1.POST("/branch-products", h.CreateBranchProduct)
	v1.GET("/branch-products", h.GetListBranchProduct)
	v1.GET("/branch-products/:id", h.GetBranchProduct)
	v1.PUT("/branch-products/:id", h.UpdateBranchProduct)
	v1.DELETE("/branch-products/:id", h.DeleteBranchProduct)

	// product api
	v1.POST("/products", h.CreateProduct)
	v1.GET("/products", h.GetListProduct)
	v1.GET("/products/:id", h.GetProduct)
	v1.PUT("/products/:id", h.UpdateProduct)
	v1.DELETE("/products/:id", h.DeleteProduct)

	// category api
	v1.POST("/categories", h.CreateCategory)
	v1.GET("/categories", h.GetListCategory)
	v1.GET("/categories/:id", h.GetCategory)
	v1.PUT("/categories/:id", h.UpdateCategory)
	v1.DELETE("/categories/:id", h.DeleteCategory)

	// staff tariff api
	v1.POST("/staff-tariffs", h.CreateStaffTariff)
	v1.GET("/staff-tariffs", h.GetListStaffTariff)
	v1.GET("/staff-tariffs/:id", h.GetStaffTariff)
	v1.PUT("/staff-tariffs/:id", h.UpdateStaffTariff)
	v1.DELETE("/staff-tariffs/:id", h.DeleteStaffTariff)

	// staff api
	v1.POST("/staffs", h.CreateStaff)
	v1.GET("/staffs", h.GetListStaff)
	v1.GET("/staffs/:id", h.GetStaff)
	v1.PUT("/staffs/:id", h.UpdateStaff)
	v1.DELETE("/staffs/:id", h.DeleteStaff)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
