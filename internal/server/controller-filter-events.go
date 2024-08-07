package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func filterEvents() gin.HandlerFunc {
	return func(context *gin.Context) {
		filter := &EventFilter{}

		if err := context.BindJSON(filter); nil != err {
			return
		}

		context.JSON(http.StatusOK, events(filter))
	}
}
