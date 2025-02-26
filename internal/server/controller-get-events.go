package server

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed json/events.json
var jsonEvents string

func getEvents() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, events(nil, context))
	}
}
