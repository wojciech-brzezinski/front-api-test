package server

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed json/profile.json
var jsonProfile string

func getProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Data(http.StatusOK, gin.MIMEJSON, []byte(jsonProfile))
	}
}
