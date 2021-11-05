package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/usecase"
)

type UserHandler struct {
	// TODO: logger
	usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{u}
}

// handle to create user
func (u *UserHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var requestBody userCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UserUseCase.CreateUser(context.Background(), requestBody.Name, requestBody.Email, requestBody.Image)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&userResponse{
		User: user,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

type userCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type userResponse struct {
	User *model.User `json:"user"`
}

func (u *UserHandler) HandleGetByProviderAccountID(w http.ResponseWriter, r *http.Request) {
	providerID := r.URL.Query().Get("provider_id")
	providerAccountID := r.URL.Query().Get("provider_account_id")

	user, err := u.UserUseCase.GetUserByProviderAccountID(context.Background(), providerID, providerAccountID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userResponse{
		User: user,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (u *UserHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	email := r.URL.Query().Get("email")

	user, err := u.UserUseCase.GetUser(context.Background(), id, email)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userResponse{
		User: user,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (u *UserHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var requestBody userUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UpdateUser(context.Background(), id, requestBody.Name, requestBody.Email, requestBody.Image)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(&userResponse{
		User: user,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

type userUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}
