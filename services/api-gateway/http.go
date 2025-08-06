package main

import "net/http"

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := readJSON(w, r, &reqBody); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, http.StatusCreated, "ok"); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}
