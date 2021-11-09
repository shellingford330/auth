package rdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/usecase/query"
)

type sessionQueryServiceImpl struct {
	DB *sql.DB
}

func NewSessionQueryService(db *sql.DB) query.SessionQueryService {
	return &sessionQueryServiceImpl{db}
}

func (s *sessionQueryServiceImpl) GetSessionBySessionToken(
	ctx context.Context,
	sessionToken string,
) (*model.Session, error) {
	session := model.Session{}
	err := s.DB.QueryRow(
		"SELECT expires, session_token, access_token, user_id FROM sessions WHERE session_token = ?",
		sessionToken,
	).Scan(&session.Expires, &session.SessionToken, &session.AccessToken, &session.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}
