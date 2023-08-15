package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

type cartIDRequest struct {
	ID int64 `uri:"id" binding:"min=1"`
}

func (server *Server) GetCartDetails(ctx *gin.Context) {
	var reqID cartIDRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		reqID.ID = -1
	}

	cart, err := server.business.Cart().GetByID(ctx, reqID.ID)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(cart))
}

func (server *Server) AddCartItem(ctx *gin.Context) {
	var req model.ModifyCartItemParams
	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	err := server.business.Cart().AddItem(ctx, &req)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(true))
}

func (server *Server) RemoveCartItem(ctx *gin.Context) {
	req := model.ModifyCartItemParams{
		Quantity: 1,
	}

	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	if err := server.business.Cart().RemoveItem(ctx, &req); err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(true))
}

func (server *Server) Checkout(ctx *gin.Context) {
	payment, err := server.business.Cart().Checkout(ctx)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(payment))
}
