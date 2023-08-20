package middleware

import (
	"log"
	"net/http"

	"github.com/hoangtk0100/dc-go-23/ex_08/pkg/util"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(util.StatusCodeCarrier); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode(), appErr)
				} else {
					// General panic cases
					ctx.AbortWithStatusJSON(
						http.StatusInternalServerError,
						util.ErrInternalServerError,
					)
				}

				log.Printf("%+v\n", err)

				if gin.IsDebugging() {
					panic(err)
				}
			}
		}()

		ctx.Next()
	}
}
