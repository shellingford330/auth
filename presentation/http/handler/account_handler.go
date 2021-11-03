package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shellingford330/auth/usecase"
)

type AccountHandler struct {
	usecase.AccountUseCase
}

func NewAccountHandler(a usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{a}
}

func (a *AccountHandler) HandleLinkAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody linkAccountRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = a.LinkAccount(context.Background(), &usecase.LinkAccountParams{
		ProviderID:        requestBody.ProviderID,
		ProviderType:      requestBody.ProviderType,
		ProviderAccountID: requestBody.ProviderAccountID,
		UserID:            requestBody.UserID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type linkAccountRequest struct {
	ProviderID        string `json:"provider_id"`
	ProviderType      string `json:"provider_type"`
	ProviderAccountID string `json:"provider_account_id"`
	UserID            string `json:"user_id"`
}
