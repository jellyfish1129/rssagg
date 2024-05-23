package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Reseponding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"` // Marshalを使う
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) { //code = http status code
	dat, err := json.Marshal(payload) // TODO: what is marshal?
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json") // TODO: what is this?
	w.WriteHeader(code)
	w.Write(dat)

}
