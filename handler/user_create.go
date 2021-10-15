package handler

import (
	"encoding/json"
	"net/http"
)

func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	var requestBody userCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// CORS対応
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

type userCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type userCreateResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}
