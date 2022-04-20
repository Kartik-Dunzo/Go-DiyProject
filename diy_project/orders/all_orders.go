package orders

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrders(c *gin.Context) {
	// Find all orders
	var Order []models.Order
	//var products []models.ProductsPurchased .Preload("CartProducts")
	if err := service.DbConn.Preload("CartProducts").Find(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request !!!"})
		return
	}
	//checking if no orders found
	if len(Order) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders Found!!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Order})
}
