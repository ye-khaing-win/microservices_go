package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/types"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	url := "http://trip-service/preview"
	log.Println(url)
	res, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to send request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		http.Error(w, "failed to send request", http.StatusInternalServerError)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "failed to read response body", http.StatusInternalServerError)
		return
	}
	var data types.OSRMApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "failed to parse response", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{
		Data: data,
	}

	log.Println("SUCCESS")
	writeJSON(w, http.StatusCreated, response)
}
