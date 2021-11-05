package rdb

import (
	"context"
	"database/sql"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
	"github.com/shellingford330/auth/pkg/ulid"
)

type sessionRepositoryImpl struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) repository.SessionRepository {
	return &sessionRepositoryImpl{db}
}

func (s *sessionRepositoryImpl) InsertSession(ctx context.Context, session *model.Session) (*model.Session, error) {
	stmt, err := s.DB.Prepare("INSERT INTO sessions (id, expires, session_token, user_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	id := ulid.Generate()
	if _, err = stmt.Exec(id, session.Expires, session.SessionToken, session.UserID); err != nil {
		return nil, err
	}
	if err := session.SetID(id); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionRepositoryImpl) UpdateSession(ctx context.Context, session *model.Session) error {
	stmt, err := s.DB.Prepare("UPDATE sessions SET expires = ?, session_token = ?, user_id = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(session.ID, session.Expires, session.SessionToken, session.UserID, session.ID)
	if err != nil {
		return err
	}
	return nil
}
