package models

import (
	"time"
)

type WaterLevelRecord struct {
	Height   float64
	Percent  float64
	DateTime time.Time
}
