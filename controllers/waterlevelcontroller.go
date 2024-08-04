package controllers

import (
	"net/http"
	"time"

	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/aPonce2001/wlmis-web-server/models"
	"github.com/aPonce2001/wlmis-web-server/websockets"
	"github.com/gin-gonic/gin"
)

type addWaterLevelRecordJSON struct {
	Height  float64 `json:"height"`
	Percent float64 `json:"percent"`
}

func AddWaterLevelRecord(context *gin.Context) {
	var newWaterLevelRecordJSON addWaterLevelRecordJSON

	if err := context.BindJSON(&newWaterLevelRecordJSON); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record := models.WaterLevelRecord{
		Height:   newWaterLevelRecordJSON.Height,
		Percent:  newWaterLevelRecordJSON.Percent,
		DateTime: time.Now(),
	}
	data.AddWaterLevelRecord(record)

	websockets.BroadcastWaterLevel(record)

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Water level record added successfully"})
}
