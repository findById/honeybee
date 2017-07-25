package v1

type Device struct {
	Id         string `json:"id"`
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"type"`
	Locate     string `json:"locate"`
	Sensor     []Sensor `json:"sensorList"`
}
