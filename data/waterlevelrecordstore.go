package data

import (
	"github.com/aPonce2001/wlmis-web-server/models"
)

var waterLevelRecordStore = []models.WaterLevelRecord{}

func AddWaterLevelRecord(waterLevelRecord models.WaterLevelRecord) {
	waterLevelRecordStore = append(waterLevelRecordStore, waterLevelRecord)
}

func GetLastWaterLevelRecord() models.WaterLevelRecord {
	return waterLevelRecordStore[len(waterLevelRecordStore)-1]
}

func GetWaterLevelRecords() []models.WaterLevelRecord {
	return waterLevelRecordStore
}

func GetLastNWaterLevelRecords(n int) []models.WaterLevelRecord {
	if n > len(waterLevelRecordStore) {
		return waterLevelRecordStore
	}

	return waterLevelRecordStore[len(waterLevelRecordStore)-n:]
}
