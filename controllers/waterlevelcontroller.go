package controllers

import (
	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/aPonce2001/wlmis-web-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	data.AddWaterLevelRecord(models.WaterLevelRecord{
		Height:   newWaterLevelRecordJSON.Height,
		Percent:  newWaterLevelRecordJSON.Percent,
		DateTime: time.Now(),
	})

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Water level record added successfully"})
}
