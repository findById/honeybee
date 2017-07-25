package api

import (
	"net/http"
	"honeybee/models/v1"
	"encoding/json"
	"log"
	"time"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"html/template"
)

func SaveDataNodeHandler(w http.ResponseWriter, r *http.Request) {
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
	if data["dataType"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'dataType' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if data["timestamp"] == "" {
		data["timestamp"] = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	}
	if data["metadata"] == "" {
		result["statusCode"] = "401"
		result["message"] = "'metadata' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	_, err = v1.FindSensorById(data["deviceId"], data["sensorId"])
	if err != nil {
		result["statusCode"] = "401"
		result["message"] = "Invalid device info"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	item := &v1.DataNode{
		DeviceId:  data["deviceId"],
		SensorId:  data["sensorId"],
		DataType:  data["dataType"],
		Timestamp: data["timestamp"],
		Metadata:  data["metadata"],
	}
	v1.Save(item)

	result["statusCode"] = "200"
	result["message"] = "success"
	b, _ := json.Marshal(result)
	w.Write(b)
}

func FindDataNodeHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})

	deviceId := r.URL.Query().Get("deviceId")
	sensorId := r.URL.Query().Get("sensorId")
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 10
	}
	if size > 1000 {
		size = 1000
	}

	if deviceId == "" {
		result["statusCode"] = "401"
		result["message"] = "'deviceId' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}
	if sensorId == "" {
		result["statusCode"] = "401"
		result["message"] = "'sensorId' must not be null"
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	res, err := v1.Find(deviceId, sensorId, size)
	if err != nil {
		result["statusCode"] = "201"
		result["message"] = err.Error()
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	result["statusCode"] = "200"
	result["message"] = "success"
	result["result"] = res
	d, _ := json.Marshal(result)
	w.Write(d)
}

func ViewReport(w http.ResponseWriter, r *http.Request) {
	d := r.Context().Value(1)
	log.Println("===", d)
	vars := mux.Vars(r)
	log.Println(vars["deviceId"])
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := make(map[string]string)
	data["deviceId"] = vars["deviceId"]

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}