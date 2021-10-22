package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (u UserHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	// ユーザID取得
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ユーザメールアドレス取得
	email := r.URL.Query().Get("email")

	// ユーザ取得
	user, err := u.UserUseCase.GetUser(context.Background(), id, email)

	// レスポンスセット
	data, err := json.Marshal(userGetResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
