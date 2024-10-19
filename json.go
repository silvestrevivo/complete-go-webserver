package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, err string) {
	if status > 499 {
		log.Println("RespondWithError: status code is 5xx error", status)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, status, errResponse{
		Error: err,
	})
}

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	dat, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal data to JSON", data)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}
