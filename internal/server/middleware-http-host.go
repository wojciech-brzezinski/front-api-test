package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func httpHost() gin.HandlerFunc {
	return func(context *gin.Context) {
		scheme := "http"

		if context.Request.TLS != nil {
			scheme = "https"
		}

		host := fmt.Sprintf("%s://%s", scheme, context.Request.Host)

		context.Set("HttpHost", host)

		context.Next()
	}
}
