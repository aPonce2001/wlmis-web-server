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
	HeightCm float64 `json:"heightCm"`
	VolumeMl float64 `json:"volumeMl"`
	Percent  float64 `json:"percent"`
}

func AddWaterLevelRecord(context *gin.Context) {
	var newWaterLevelRecordJSON addWaterLevelRecordJSON

	if err := context.BindJSON(&newWaterLevelRecordJSON); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record := models.WaterLevelRecord{
		HeightCm: newWaterLevelRecordJSON.HeightCm,
		Percent:  newWaterLevelRecordJSON.Percent,
		VolumeMl: newWaterLevelRecordJSON.VolumeMl,
		DateTime: time.Now(),
	}
	data.AddWaterLevelRecord(record)

	websockets.BroadcastWaterLevel(record)

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Water level record added successfully"})
}
