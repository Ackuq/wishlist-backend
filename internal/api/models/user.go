package models

type CreateUser struct {
	Id       int
	Name     string
	Email    string
	Password []byte
}
