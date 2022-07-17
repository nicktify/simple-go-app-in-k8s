package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time time.Time `json:"time"`
	Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handsOn := HandsOn{
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	json.NewEncoder(w).Encode(handsOn)
}

func main () {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
