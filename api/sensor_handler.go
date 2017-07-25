package api

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"honeybee/models/v1"
)

func SaveSensorHandler(w http.ResponseWriter, r *http.Request) {
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
	if data["sensorId"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'sensorId' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if data["sensorName"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'sensorName' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if data["sensorType"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'sensorType' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	item := &v1.Sensor{
		DeviceId:   data["deviceId"],
		SensorId:   data["sensorId"],
		SensorName: data["sensorName"],
		SensorType: data["sensorType"],
	}
	v1.SaveSensor(item)

	result["statusCode"] = "200"
	result["message"] = "success"
	b, _ := json.Marshal(result)
	w.Write(b)
}

func FindSensorHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})

	deviceId := r.URL.Query().Get("deviceId")
	sensorId := r.URL.Query().Get("sensorId")
	if deviceId == "" {
		result["statusCode"] = "401"
		result["message"] = "'deviceId' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	var list []v1.Sensor
	var err error
	if sensorId == "" {
		list, err = v1.FindSensorByDeviceId(deviceId)
	} else {
		list, err = v1.FindSensorById(deviceId, sensorId)
	}
	if err != nil {
		result["statusCode"] = "201"
		result["message"] = err.Error()
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	result["statusCode"] = "200"
	result["message"] = "success"
	result["result"] = list
	d, _ := json.Marshal(result)
	w.Write(d)
}
