package models

import (
	"time"
)

type WaterLevelRecord struct {
	HeightCm float64
	Percent  float64
	VolumeMl float64
	DateTime time.Time
}
