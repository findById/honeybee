package v1

import (
	"testing"
	"log"
	"encoding/json"
)

func TestSave(t *testing.T) {
	item := &DataNode{
		DeviceId:  "deviceId",
		SensorId:  "sensorId",
		DataType:  "sht",
		Timestamp: "timestamp",
		Metadata:  "metadata",
	}
	Save(item)
}

func TestFind(t *testing.T) {
	data, err := Find("deviceId", "0", 10)
	if err != nil {
		log.Println(err)
	}
	d,err := json.Marshal(data)
	log.Println(string(d))
}