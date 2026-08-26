package repository

import (
	"database/sql"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user model.RegisterRequest) (model.UserResponse, error) {
	var result model.UserResponse
	tx, err := u.db.Beginx()
	if err != nil {
		return result, err
	}

	defer tx.Rollback()

	query := "INSERT INTO users (name, email, password, phone_number) VALUES ($1, $2, $3, $4) RETURNING name, email, phone_number, role, created_at, updated_at"
	err = tx.Get(&result, query, user.Name, user.Email, user.Password, user.PhoneNumber)

	if err != nil {
		return result, err
	}

	tx.Commit()
	return result, nil
}

func (u *UserRepository) GetUsers() ([]model.UserResponse, error) {
	var result []model.UserResponse
	err := u.db.Select(&result, "SELECT name, email, phone_number, role, created_at, updated_at FROM users")

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.QueryRowx("SELECT * FROM users WHERE email = $1", email).StructScan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}
