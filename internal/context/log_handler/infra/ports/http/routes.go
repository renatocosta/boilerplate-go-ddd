package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(
	r *mux.Router,
	controller HttpServer) {
	r.Handle("/select-log-file", ValidateRequest(http.HandlerFunc(controller.SelectLogFile))).Methods("POST")
	r.HandleFunc("/available-log-files", controller.AvailableLogFiles).Methods("GET")
}
