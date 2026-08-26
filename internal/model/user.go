package model

import "time"

type User struct {
	Id          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Role        string    `db:"role" json:"role"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type RegisterRequest struct {
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
}

type UserResponse struct {
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber *string   `db:"phone_number" json:"phone_number"`
	Role        string    `db:"role" json:"role"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type LoginRequest struct {
	Email    string `db:"email" json:"email" binding:"required"`
	Password string `db:"password" json:"password" binding:"required"`
}
