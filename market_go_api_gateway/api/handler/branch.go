package handler

import (
	"genproto/branch_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @Router       /v1/branches [post]
// @Summary      Create a new branch
// @Description  Create a new branch with the provided details
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        branch     body  branch_service.BranchCreateReq  true  "data of the branch"
// @Success      201  {object}  branch_service.BranchCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateBranch(ctx *gin.Context) {
	var branch = branch_service.Branch{}

	err := ctx.ShouldBindJSON(&branch)
	if err != nil {
		h.handlerResponse(ctx, "CreateBranch", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchService().Create(ctx, &branch_service.BranchCreateReq{
		Name:      branch.Name,
		Address:   branch.Address,
		FoundedAt: branch.FoundedAt,
	})

	if err != nil {
		h.handlerResponse(ctx, "BranchService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create branch response", http.StatusOK, resp)
}

// ListBranches godoc
// @Router       /v1/branches [get]
// @Summary      List branches
// @Description  Get branches
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        search     query     string false "search by name and address"
// @Success      200  {array}   branch_service.Branch
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListBranch(ctx *gin.Context) {
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

	resp, err := h.services.BranchService().GetList(ctx.Request.Context(), &branch_service.BranchGetListReq{
		Page:   int64(page),
		Limit:  int64(limit),
		Search: ctx.Query("search"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListBranch", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list branch response", http.StatusOK, resp)
}

// GetBranch godoc
// @Router       /v1/branches/{id} [get]
// @Summary      Get a branch by ID
// @Description  Retrieve a branch by its unique identifier
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Branch ID to retrieve"
// @Success      200  {object}  branch_service.Branch
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetBranch(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.BranchService().GetById(ctx.Request.Context(), &branch_service.BranchIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branch GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get branch response", http.StatusOK, resp)
}

// UpdateBranch godoc
// @Router       /v1/branches/{id} [put]
// @Summary      Update an existing branch
// @Description  Update an existing branch with the provided details
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Branch ID to update"
// @Param        branch   body    branch_service.BranchUpdateReq  true    "Updated data for the branch"
// @Success      200  {object}  branch_service.BranchUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateBranch(ctx *gin.Context) {
	var branch = branch_service.Branch{}
	branch.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&branch)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchService().Update(ctx.Request.Context(), &branch_service.BranchUpdateReq{
		Id:        branch.Id,
		Name:      branch.Name,
		Address:   branch.Address,
		FoundedAt: branch.FoundedAt,
	})

	if err != nil {
		h.handlerResponse(ctx, "error branch Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update branch response", http.StatusOK, resp.Msg)
}

// DeleteBranch godoc
// @Router       /v1/branches/{id} [delete]
// @Summary      Delete a branch
// @Description  Delete a branch by its unique identifier
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Branch ID to retrieve"
// @Success      200  {object}  branch_service.BranchDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteBranch(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.BranchService().Delete(ctx.Request.Context(), &branch_service.BranchIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branch Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete branch response", http.StatusOK, resp.Msg)
}
