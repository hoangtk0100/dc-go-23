package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, err error) {
	if cErr, ok := err.(StatusCodeCarrier); ok {
		ctx.JSON(cErr.StatusCode(), cErr)
		return
	}

	ctx.JSON(
		http.StatusInternalServerError,
		ErrInternalServerError.WithError(err.Error()),
	)
}
