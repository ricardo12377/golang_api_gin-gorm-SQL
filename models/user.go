package models

import "time"

type User struct {
	ID uint `json: "id" gorm: "primary_key"`
	Name string `json: "name"`
	Password string `json: "password`
	Email string `json: "email"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	DeletedAt time.Time `json: "deleted_at" gorm: "index"`
}