package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mainRoutes "github.com/lealhugui/vice/routeManager"
	"github.com/prometheus/common/log"
)

func StartServer() {
	//r.Use(static.Serve("/", static.LocalFile("./static/", false)))

	r := gin.Default()
	r.POST("/registerTask", mainRoutes.RegisterTask)
	r.POST("/registerWorker", mainRoutes.RegisterWorker)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := 8080

	go func() {
		log.Error(r.Run(fmt.Sprintf(":%d", port)))
	}()

	log.Info(fmt.Sprintf("Server Started on Port:%d", port))
}
