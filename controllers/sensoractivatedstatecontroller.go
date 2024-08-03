package controllers

import (
	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToggleSensorActivatedState(c *gin.Context) {
	sensorActivated := data.ToggleSensorActivatedState()
	message := "Sensor deactivated"
	if sensorActivated {
		message = "Sensor activated"
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}
