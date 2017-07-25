package api

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"honeybee/models/v1"
)

func SaveDeviceHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result["statusCode"] = "201"
		result["message"] = err.Error()
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	log.Println(string(body))
	var data map[string]string
	json.Unmarshal(body, &data)
	if data["deviceId"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'deviceId' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if data["deviceType"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'deviceType' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if data["locate"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'locate' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	item := &v1.Device{
		DeviceId:  data["deviceId"],
		DeviceType:  data["deviceType"],
		Locate:  data["locate"],
	}
	v1.SaveDevice(item)

	result["statusCode"] = "200"
	result["message"] = "success"
	b, _ := json.Marshal(result)
	w.Write(b)
}

func FindDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}