package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shellingford330/auth/usecase"
)

type SessionHandler struct {
	usecase.SessionUseCase
}

func NewSessionHandler(s usecase.SessionUseCase) *SessionHandler {
	return &SessionHandler{s}
}

func (s *SessionHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var requestBody createSessionRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := s.CreateSession(context.Background(), &usecase.CreateSessionParams{
		Expires:      requestBody.Expires,
		SessionToken: requestBody.SessionToken,
		UserID:       requestBody.UserID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&sessionResponse{
		Expires:      session.Expires,
		SessionToken: session.SessionToken,
		UserID:       session.UserID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

type createSessionRequest struct {
	Expires      time.Time `json:"expires"`
	SessionToken string    `json:"session_token"`
	UserID       string    `json:"user_id"`
}

type sessionResponse struct {
	Expires      time.Time `json:"expires"`
	SessionToken string    `json:"session_token"`
	UserID       string    `json:"user_id"`
}

func (s *SessionHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.URL.Query().Get("session_token")

	session, err := s.SessionUseCase.GetSession(context.Background(), sessionToken)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&sessionResponse{
		Expires:      session.Expires,
		SessionToken: session.SessionToken,
		UserID:       session.UserID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
