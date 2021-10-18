package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shellingford330/auth/domain/model"
)

func (u UserHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	// リクエストデコード
	var requestBody userCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// ユーザ登録
	user := &model.User{
		Name:  requestBody.Name,
		Email: requestBody.Email,
		Image: requestBody.Image,
	}
	user, err = u.UserRepository.InsertUser(context.Background(), user)
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
