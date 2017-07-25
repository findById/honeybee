package v1

type Sensor struct {
	Id         string `json:"id"`
	DeviceId   string `json:"deviceId"`
	SensorId   string `json:"sensorId"`
	SensorName string `json:"name"`
	SensorType string `json:"type"`
	Data       []DataNode `json:"dataList"`
}
