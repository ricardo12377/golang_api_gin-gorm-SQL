package dtos

type UpdateUserInput struct {
	Name string `json: "name"`
	Password string `json: "password"`
	Email string `json: "email"`
}