package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(port uint16) error {
	router := gin.Default()

	router.Use(delay())

	router.NoRoute(notFound())

	router.Static("/images", "/images")

	router.GET("/events", getEvents())
	router.POST("/events", filterEvents())
	router.GET("/profile", getProfile())

	return router.Run(
		fmt.Sprintf(":%d", port),
	)
}
