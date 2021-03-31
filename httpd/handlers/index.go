package handlers

import "net/http"

// Index is the home of the API
func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{
		"message": "Ryan Back Raspberry A PI",
	}
	respondWithJSON(w, http.StatusOK, response)
}
