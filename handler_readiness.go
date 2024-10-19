package main

import "net/http"

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, struct{}{})
}
