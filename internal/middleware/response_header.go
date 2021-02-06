package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-backend-template/global"
)

func ResponseHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		// add customize response header
		for header, value := range global.HTTPServerSetting.ResponseHeaders {
			ctx.Header(header, value)
		}
	}
}
