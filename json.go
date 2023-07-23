package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respond with 5XX error: ", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}