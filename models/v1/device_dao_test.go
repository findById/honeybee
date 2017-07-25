package v1

import (
	"testing"
	"encoding/json"
	"log"
)

func TestSaveDevice(t *testing.T) {
	item := &Device{
		DeviceId:   "deviceId",
		DeviceType: "ESP8266",
		Locate:     "",
	}
	SaveDevice(item)
}

func TestFindDeviceById(t *testing.T) {
	data, err := FindDeviceById("deviceId")
	if err != nil {
		log.Println(err)
	}
	d, err := json.Marshal(data)
	log.Println(string(d))
}
