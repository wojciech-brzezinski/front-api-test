package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func notFound() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusNotFound, "")
	}
}
