package models

import (
	"diy_project/service"
	"fmt"
	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func CreateModels() {
	Conn = service.DbConn
	Conn.DropTable(&Products{})
	Conn.DropTable(&Users{})
	Conn.DropTable(&Order{})
	Conn.DropTable(&ProductsPurchased{})
	Conn.AutoMigrate(&Products{}, &Users{}, &Order{}, &ProductsPurchased{})
	fmt.Println("Models Product, User and Order has created successfully!!!")
}
