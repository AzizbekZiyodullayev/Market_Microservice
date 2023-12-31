package handler

import (
	"genproto/sale_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSaleProduct godoc
// @Router       /v1/sale-products [post]
// @Summary      Create a new sale-product
// @Description  Create a new sale-product with the provided details
// @Tags         Sale-products
// @Accept       json
// @Produce      json
// @Param        sale-product     body  sale_service.SaleProductCreateReq  true  "data of the sale-product"
// @Success      201  {object}  sale_service.SaleProductCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateSaleProduct(ctx *gin.Context) {
	var saleProduct = sale_service.SaleProduct{}

	err := ctx.ShouldBindJSON(&saleProduct)
	if err != nil {
		h.handlerResponse(ctx, "CreateSaleProduct", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Create(ctx, &sale_service.SaleProductCreateReq{
		SaleId:    saleProduct.SaleId,
		ProductId: saleProduct.ProductId,
		Price:     saleProduct.Price,
		Quantity:  saleProduct.Quantity,
	})

	if err != nil {
		h.handlerResponse(ctx, "SaleProductService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create sale product response", http.StatusOK, resp)
}

// ListSaleProducts godoc
// @Router       /v1/sale-products [get]
// @Summary      List sale-products
// @Description  Get sale-products
// @Tags         Sale-products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Success      200  {array}   sale_service.SaleProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListSaleProduct(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		h.handlerResponse(ctx, "error get page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		h.handlerResponse(ctx, "error get limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().GetList(ctx.Request.Context(), &sale_service.SaleProductGetListReq{
		Page:  int64(page),
		Limit: int64(limit),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListSaleProduct", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list sale product response", http.StatusOK, resp)
}

// GetSaleProduct godoc
// @Router       /v1/sale-products/{id} [get]
// @Summary      Get a sale_product by ID
// @Description  Retrieve a sale by its unique identifier
// @Tags         Sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id   path    string     true    "SaleProduct ID to retrieve"
// @Success      200  {object}  sale_service.SaleProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetSaleProduct(ctx *gin.Context) {
	saleId := ctx.Param("sale_id")

	resp, err := h.services.SaleProductService().GetById(ctx.Request.Context(), &sale_service.SaleProductIdReq{
		SaleId: saleId,
	})
	if err != nil {
		h.handlerResponse(ctx, "error sale GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get sale product response", http.StatusOK, resp)
}

// UpdateSaleProduct godoc
// @Router       /v1/sale-products/{id} [put]
// @Summary      Update an existing sale_product
// @Description  Update an existing sale_product with the provided details
// @Tags         Sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id       path    string     true    "sale ID to update"
// @Param        sale_product   body    sale_service.SaleProductUpdateReq  true    "Updated data for the sale_product"
// @Success      200  {object}  sale_service.SaleProductUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateSaleProduct(ctx *gin.Context) {
	var saleProduct = sale_service.SaleProduct{}
	saleProduct.SaleId = ctx.Param("sale_id")
	if err := ctx.ShouldBindJSON(&saleProduct); err != nil {
		h.handlerResponse(ctx, "error update sale product", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Update(ctx.Request.Context(), &sale_service.SaleProductUpdateReq{
		SaleId:    saleProduct.SaleId,
		ProductId: saleProduct.ProductId,
		Quantity:  saleProduct.Quantity,
		Price:     saleProduct.Price,
	})

	if err != nil {
		h.handlerResponse(ctx, "error sale Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update sale product response", http.StatusOK, resp.Msg)
}

// DeleteSaleProduct godoc
// @Router       /v1/sale-products/{id} [delete]
// @Summary      Delete a sale-product
// @Description  Delete a sale-product by its unique identifier
// @Tags         Sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.SaleProductDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteSaleProduct(ctx *gin.Context) {
	saleId := ctx.Param("sale_id")

	resp, err := h.services.SaleProductService().Delete(ctx.Request.Context(), &sale_service.SaleProductIdReq{
		SaleId: saleId,
	})
	if err != nil {
		h.handlerResponse(ctx, "error sale product Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete sale product response", http.StatusOK, resp.Msg)
}
