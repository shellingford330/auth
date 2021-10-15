package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func (u UserHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	// リクエストデコード
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// ユーザ登録
	user, err := u.UserRepository.GetUser(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// レスポンスセット
	data, err := json.Marshal(&userGetResponse{
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

type userGetResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}
