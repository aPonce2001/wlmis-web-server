package controllers

import (
	"net/http"

	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/aPonce2001/wlmis-web-server/websockets"
	"github.com/gin-gonic/gin"
)

func ToggleSensorActivatedState(c *gin.Context) {
	sensorActivated := data.ToggleSensorActivatedState()
	message := "Sensor deactivated"
	if sensorActivated {
		message = "Sensor activated"
	}

	websockets.BroadcastSensorActivatedState(sensorActivated)
	c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}
