package server

import (
	"diy_project/config"
	"diy_project/orders"
	"diy_project/product"
	"diy_project/user"
	"github.com/gin-gonic/gin"
)

var configure *config.Config

func StartServer() {
	configure = &config.Config_parse

	router := gin.Default()
	//users
	router.POST("user/create", user.NewUser)
	router.PATCH("user/:id", user.UpdateUser)
	//products
	router.POST("add_products_list", product.AddLIstOfProducts)
	router.GET("product/:product_id", product.GetProductsById)
	router.GET("all_products", product.GetAllProducts)
	router.PATCH("product/:product_id", product.UpdateProducts)
	//orders
	router.POST("place_order", orders.BuyProducts)
	router.GET("top_products", orders.GetTopPProducts)
	router.GET("all_orders", orders.GetAllOrders)
	//start server
	router.Run(configure.Server.IP + ":" + configure.Server.Port)

}
