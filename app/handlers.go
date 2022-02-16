package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(rw http.ResponseWriter, r *http.Request) {
	response := make(map[string]time.Time)
	params := r.URL.Query()
	timezone := params.Get("tz")
	if timezone != "" {
		for _, t := range strings.Split(timezone, ",") {
			loc, err := time.LoadLocation(t)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				fmt.Fprint(rw, "Invalid timezone.")
				json.NewEncoder(rw)
				log.Printf("Timezone %s is invalid.", t)
				return
			} else {
				response[t] = time.Now().In(loc)
			}
		}
	} else {
		response["current_time"] = time.Now()
	}
	json.NewEncoder(rw).Encode(response)
}
