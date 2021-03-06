package model

import (
	"errors"
	"time"
)

type Session struct {
	Expires      time.Time `json:"expires"`
	SessionToken string    `json:"session_token"`
	AccessToken  string    `json:"access_token"`
	UserID       string    `json:"user_id"`
}

func NewSession(expires time.Time, userID string) (*Session, error) {
	session := Session{}
	if err := session.SetExpires(expires); err != nil {
		return nil, err
	}
	if err := session.SetUserID(userID); err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *Session) SetExpires(expires time.Time) error {
	if expires.IsZero() {
		return errors.New("expires is blank")
	}
	s.Expires = expires
	return nil
}

func (s *Session) SetSessionToken(sessionToken string) error {
	if sessionToken == "" {
		return errors.New("sessionToken is blank")
	}
	s.SessionToken = sessionToken
	return nil
}

func (s *Session) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return errors.New("accessToken is blank")
	}
	s.AccessToken = accessToken
	return nil
}

func (s *Session) SetUserID(userID string) error {
	if userID == "" {
		return errors.New("userID is blank")
	}
	s.UserID = userID
	return nil
}
