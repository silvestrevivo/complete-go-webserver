package main

import "net/http"

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 400, "Something went wrong")
}
