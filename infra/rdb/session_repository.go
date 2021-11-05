package rdb

import (
	"context"
	"database/sql"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type sessionRepositoryImpl struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) repository.SessionRepository {
	return &sessionRepositoryImpl{db}
}

func (s *sessionRepositoryImpl) InsertSession(ctx context.Context, session *model.Session) (*model.Session, error) {
	stmt, err := s.DB.Prepare("INSERT INTO sessions (expires, session_token, user_id) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	if _, err = stmt.Exec(session.Expires, session.SessionToken, session.UserID); err != nil {
		return nil, err
	}
	return session, nil
}
