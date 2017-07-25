package v1

import (
	"testing"
	"encoding/json"
	"log"
)

func TestSaveSensor(t *testing.T) {
	item := &Sensor{
		DeviceId:"deviceId",
		SensorId:"0",
		SensorName:"SENSOR",
		SensorType:"SHT",
	}
	SaveSensor(item)
}

func TestFindSensorById(t *testing.T) {
	data, err := FindSensorById("deviceId", "0")
	if err != nil {
		log.Println(err)
	}
	d, err := json.Marshal(data)
	log.Println(string(d))
}
