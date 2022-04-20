//go:generate mockgen -source product_struct.go -destination mock/product_struct_mock.go -package mock

package models

type ProductInterface interface {
	ListOfProducts(r ListOfProducts) error
	Products(r Products) error
	UpdateProductInput(r UpdateProductInput) error
}

type ListOfProducts struct {
	Product_list []Products `json:"product_list"`
	User_Id      int        `gorm:"not null"json:"user_id"`
}
type Products struct {
	Id            uint   `gorm:"primary_key;AUTO_INCREMENT"json:"id"`
	User_Id       int    `gorm:"not null"json:"user_id"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Price         int    `json:"price"`
	Quantity      int    `json:"quantity"`
	Sold_Quantity int    `json:"sold_quantity"`
	Status        string `gorm:"DEFAULT:'active'"json:"status"`
}

type UpdateProductInput struct {
	Name          string `json:"name"`
	Category      string `json:"category"`
	Price         int    `json:"price"`
	Quantity      int    `json:"quantity"`
	Sold_Quantity int    `json:"sold_quantity"`
	Status        string `json:"status"`
}
