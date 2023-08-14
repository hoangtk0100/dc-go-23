package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
)

type productIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// CreateProduct create a new product
func (server *Server) CreateProduct(ctx *gin.Context) {
	var req model.CreateProductParams
	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	prod, err := server.business.Product().Create(ctx, &req)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(prod))
}

// UpdateProduct update a product
func (server *Server) UpdateProduct(ctx *gin.Context) {
	var reqID productIDRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	var req model.UpdateProductParams
	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	req.ID = reqID.ID
	prod, err := server.business.Product().Update(ctx, &req)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(prod))
}

// DeleteProductByID delete a product by ID
func (server *Server) DeleteProductByID(ctx *gin.Context) {
	var reqID productIDRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	err := server.business.Product().DeleteByID(ctx, reqID.ID)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(true))
}

// GetProductByID get a product by ID
func (server *Server) GetProductByID(ctx *gin.Context) {
	var reqID productIDRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	prod, err := server.business.Product().GetByID(ctx, reqID.ID)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(prod))
}

// GetProducts get all products
func (server *Server) GetProducts(ctx *gin.Context) {
	prods, err := server.business.Product().GetAll(ctx)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(prods))
}
