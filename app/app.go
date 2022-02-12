package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() {
	route := mux.NewRouter()

	route.HandleFunc("/api/time", getCurrentTime).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8080", route)
}

