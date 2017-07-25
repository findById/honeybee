package main

import (
	"flag"
	"github.com/gorilla/mux"
	"honeybee/api"
	"net/http"
)

var (
	port = flag.String("port", "8080", "accept port")
)

func main() {
	flag.Parse()
	if *port == "" {
		flag.PrintDefaults()
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/view/{deviceId}", api.ViewReport).Methods("GET")

	r.HandleFunc("/device/save", api.SaveDeviceHandler).Methods("POST")
	r.HandleFunc("/device/find", api.FindDeviceHandler).Methods("GET")

	r.HandleFunc("/sensor/save", api.SaveSensorHandler).Methods("POST")
	r.HandleFunc("/sensor/find", api.FindSensorHandler).Methods("GET")

	r.HandleFunc("/sensor/data/save", api.SaveDataNodeHandler).Methods("POST")
	r.HandleFunc("/sensor/data/find", api.FindDataNodeHandler).Methods("GET")

	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":" + *port, nil)
}
