package api

import (
	"testing"
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
	"time"
	"fmt"
	"math/rand"
)

func TestSaveDataNodeHandler(t *testing.T) {
	req := `{"deviceId": "deviceId","sensorId": "sensorId","dataType": "0","timestamp": "","metadata": "metadata"}`

	resp, err := http.Post("http://localhost:8080/sensor/data/save", "application/json", bytes.NewReader([]byte(req)))
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}

func TestFindDataNodeHandler(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/sensor/data/find?deviceId=deviceId&sensorId=0")
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}

func TestSaveDataTask(t *testing.T) {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for range ticker.C {
			req := fmt.Sprintf(`{"deviceId": "deviceId","sensorId": "0","dataType": "0","timestamp": "","metadata": "%v"}`,
				rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100))
			resp, err := http.Post("http://localhost:8080/sensor/data/save", "application/json", bytes.NewReader([]byte(req)))
			if err != nil {
				log.Println(err)
				break
			}
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(string(b))
		}
	}()
	select {}
}
