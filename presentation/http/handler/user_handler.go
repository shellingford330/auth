package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/shellingford330/auth/usecase"
)

type UserHandler struct {
	// TODO: logger
	usecase.UserUseCase
}

// handle to create user
func (u UserHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var requestBody userCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		httpError(w, err, http.StatusBadRequest)
	}

	user, err := u.UserUseCase.CreateUser(context.Background(), requestBody.Name, requestBody.Email, requestBody.Image)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}

	data, err := json.Marshal(&userCreateResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}
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

func (u UserHandler) HandleGetByProviderAccountID(w http.ResponseWriter, r *http.Request) {
	// ユーザメールアドレス取得
	providerID := r.URL.Query().Get("provider_id")
	providerAccountID := r.URL.Query().Get("provider_account_id")

	user, err := u.UserUseCase.GetUserByProviderAccountID(context.Background(), providerID, providerAccountID)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}

	data, err := json.Marshal(userGetResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}
	w.Write(data)
}

func (u UserHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		httpError(w, err, http.StatusBadRequest)
	}
	email := r.URL.Query().Get("email")

	user, err := u.UserUseCase.GetUser(context.Background(), id, email)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}

	data, err := json.Marshal(userGetResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
	}
	w.Write(data)
}

type userGetResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

func httpError(w http.ResponseWriter, err error, code int) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
