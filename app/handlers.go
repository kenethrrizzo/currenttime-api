package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(rw http.ResponseWriter, r *http.Request) {
	response := make(map[string]time.Time)
	params := r.URL.Query()
	timezone := params.Get("tz")
	for _, t := range strings.Split(timezone, ",") {
		loc, err := time.LoadLocation(t)
		if err != nil {
			rw.WriteHeader(404)
			log.Fatalf("Error finding timezone. Err: %s", err) // corregir
		}
		response[t] = time.Now().In(loc)
	}
	json.NewEncoder(rw).Encode(response)
}