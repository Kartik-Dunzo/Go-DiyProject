//go:generate mockgen -source user_struct.go -destination mock/user_struct_mock.go -package mock Users UpdateUserInput

package models

type UserInterface interface {
	Users(r Users) error
	UpdateUserInput(r UpdateUserInput) error
}

type Users struct {
	Id     uint   `gorm:"primary_key;AUTO_INCREMENT"json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Type   string `json:"type"`
	Status string `gorm:"DEFAULT:'active'"json:"status"`
}

type UpdateUserInput struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Type   string `json:"type"`
	Status string `json:"status"`
}
