package main

import (
	"./points"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// API ROUTES

	router.HandleFunc("/points/add", points.Add).Methods("POST")
	router.HandleFunc("/points/deduct", points.Deduct).Methods("POST")
	router.HandleFunc("/points/points", points.Get).Methods("GET")

	// start server
	glog.Info("Starting Points api on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		glog.Fatal("ListenAndServe: ", err)
	}
}
