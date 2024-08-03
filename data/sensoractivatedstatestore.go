package data

var sensorActivatedState = false

func GetSensorActivatedState() bool {
	return sensorActivatedState
}

func ToggleSensorActivatedState() bool {
	sensorActivatedState = !sensorActivatedState

	return sensorActivatedState
}
