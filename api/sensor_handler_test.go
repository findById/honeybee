package api

import (
	"testing"
	"bytes"
	"net/http"
	"log"
	"io/ioutil"
)

func TestSaveSensorHandler(t *testing.T) {
	req := `{"deviceId": "deviceId","sensorId":"5","sensorName":"SENSOR","sensorType": "HALL"}`

	resp, err := http.Post("http://localhost:8080/sensor/save", "application/json", bytes.NewReader([]byte(req)))
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}

func TestFindSensorHandler(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/sensor/find?deviceId=deviceId&sensorId=0")
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}