package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

func (u UserHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	// リクエストデコード
	var requestBody userCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// ユーザ登録
	user, err := u.UserUseCase.CreateUser(context.Background(), requestBody.Name, requestBody.Email, requestBody.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// レスポンスセット
	data, err := json.Marshal(userCreateResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
