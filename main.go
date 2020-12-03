package main

import (
	"PointsCalculator/points"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	router := mux.NewRouter()

	// API ROUTES

	router.HandleFunc("/points/add", points.Add).Methods("POST")
	router.HandleFunc("/points/deduct", points.Deduct).Methods("POST")
	router.HandleFunc("/points/balance", points.Balance).Methods("GET")

	// start server
	glog.Info("Starting Points api on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		glog.Fatal("ListenAndServe: ", err)
	}
}
