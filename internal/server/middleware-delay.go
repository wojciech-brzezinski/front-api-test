package server

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func delay() gin.HandlerFunc {
	return func(context *gin.Context) {
		ms := rand.Intn(2000) + 1

		time.Sleep(time.Duration(ms) * time.Millisecond)

		context.Next()
	}
}
