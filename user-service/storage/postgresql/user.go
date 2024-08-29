package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"user-service/pkg/hash"
	userpb "user-service/protos/user"

	"github.com/lib/pq"
)

type Storage interface {
	InsertUser(ctx context.Context, req *userpb.RegisterUserRequest) (int, error)
	LoginSql(ctx context.Context, req *userpb.LoginUserRequest) (bool, error)
	CheckUserIDSql(ctx context.Context, req string) (bool, error)
	GetUserSql(ctx context.Context, userID string) (*userpb.GetUserResponse, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) InsertUser(ctx context.Context, req *userpb.RegisterUserRequest) (int, error) {
	var userID int
	err := s.db.QueryRowContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", req.Username, req.Email, req.Password).Scan(&userID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return 0, errors.New("username already exists")
			}
		}
		return 0, err
	}
	return userID, nil
}

func (s *PostgresStorage) LoginSql(ctx context.Context, req *userpb.LoginUserRequest) (bool, error) {
	var storedPassword string
	err := s.db.QueryRowContext(ctx, "SELECT password FROM users WHERE username = $1", req.Username).Scan(&storedPassword)
	if err != nil {
		return false, errors.New("invalid username")
	}

	check := hash.CheckPasswordHash(req.Password, storedPassword)

	if !check {
		return false, errors.New("invalid password")
	}
	return true, nil
}

func (s *PostgresStorage) CheckUserIDSql(ctx context.Context, userID string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *PostgresStorage) GetUserSql(ctx context.Context, userID string) (*userpb.GetUserResponse, error) {
	var user userpb.GetUserResponse
	err := s.db.QueryRowContext(ctx, "SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.UserID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
