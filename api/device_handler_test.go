package api

import (
	"testing"
	"bytes"
	"net/http"
	"log"
	"io/ioutil"
)

func TestSaveDeviceHandler(t *testing.T) {
	req := `{"deviceId": "deviceId","deviceType": "deviceType","locate": "locate"}`

	resp, err := http.Post("http://localhost:8080/device/save", "application/json", bytes.NewReader([]byte(req)))
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}

func TestFindDeviceHandler(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/device/find?deviceId=deviceId")
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}