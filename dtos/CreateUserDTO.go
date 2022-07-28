package dtos

type CreateUserInput struct {
 	Name string `json: "name"`
	Password string `json: "password`
	Email string `json: "email"`
}