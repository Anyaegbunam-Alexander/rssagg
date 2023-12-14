package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	// since we only care about the 200 use an empty struct for r
	respondWithError(w, 400, "Something went wrong")
}
