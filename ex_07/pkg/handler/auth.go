package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

func (server *Server) Login(ctx *gin.Context) {
	var req model.LoginParams
	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	tokens, err := server.business.Auth().Login(ctx, &req)
	if err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(tokens))
}
