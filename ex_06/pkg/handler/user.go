package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
)

func (server *Server) Register(ctx *gin.Context) {
	var req model.CreateUserParams
	if err := ctx.ShouldBind(&req); err != nil {
		util.ErrorResponse(ctx, util.ErrBadRequest.WithError(err.Error()))
		return
	}

	if _, err := server.business.User().Register(ctx, &req); err != nil {
		util.ErrorResponse(ctx, err)
		return
	}

	util.SuccessResponse(ctx, util.NewDataResponse(true))
}
