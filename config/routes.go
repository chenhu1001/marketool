// routes.go

package config

import (
	"github.com/chenhu1001/marketool/handlers"
	"github.com/chenhu1001/marketool/internal"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	// GET /
	//router.GET("/", showIndexPage)
	internal.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	internal.Router.GET("/", handlers.ShowIndexPage)
	internal.Router.GET("/houses", handlers.GetHomeHouse)
	internal.Router.GET("/houses/:house_id", handlers.GetHouse)
}
