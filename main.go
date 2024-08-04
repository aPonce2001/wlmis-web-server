package main

import (
	"github.com/aPonce2001/wlmis-web-server/routes"
	"github.com/aPonce2001/wlmis-web-server/websockets"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.ConfigureRoutes(router)
	websockets.StartWaterLevelWebSocket(router)
	websockets.StartSensorWebSocket(router)
	router.Run("0.0.0.0:5000")
}
