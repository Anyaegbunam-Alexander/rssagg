package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// since we only care about the 200 use an empty struct for r
	respondWithJSON(w, 200, struct{}{})
}
