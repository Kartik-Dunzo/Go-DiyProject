//go:generate mockgen -source order_struct.go -destination mock/order_struct_mock.go -package mock
package models

import (
	"time"
)

type OrderInterface interface {
	ProductsPurchased(r ProductsPurchased) error
	Order(r Order) error
}

type ProductsPurchased struct {
	OrderId         int `json:"order_id"`
	ProductID       int `json:"product_id"`
	ProductQuantity int `json:"product_quantity"`
}
type Order struct {
	Id           int                 `gorm:"primary_key;AUTO_INCREMENT"json:"id"`
	UserId       int                 `json:"user_id"`
	CartProducts []ProductsPurchased `gorm:"foreignKey:OrderId;references:Id"json:"cart_products"`
	CreatedAt    time.Time           `json:"created_at"`
}
