package main

import (
	"github.com/aPonce2001/wlmis-web-server/routes"
	"github.com/aPonce2001/wlmis-web-server/websockets"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.ConfigureRoutes(router)
	websockets.StartWaterLevelWebSocket(router)
	websockets.StartSensorWebSocket(router)
	router.Run(":5000")
}
