package routes

import (
	"github.com/aPonce2001/wlmis-web-server/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) {
	router.POST("/api/water-level", controllers.AddWaterLevelRecord)
	router.POST("/api/sensor-activated-state", controllers.ToggleSensorActivatedState)
}
