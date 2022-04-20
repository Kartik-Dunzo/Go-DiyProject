package orders

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GroupProducts struct {
	ProductID int `json:"product_id"`
	Total     int `json:"total"`
}

func GetTopPProducts(c *gin.Context) {
	// Check Product exist
	var OrderId []int
	//var products []models.ProductsPurchased .Preload("CartProducts")
	if err := service.DbConn.Model(&models.Order{}).Where("created_at >= ?", time.Now().Add(-time.Hour*12)).Pluck("id", &OrderId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request !!!"})
		return
	}
	//checking if no orders found
	if len(OrderId) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Products sold in the last hour!!!"})
		return
	}
	var TopSellingProducts []GroupProducts
	service.DbConn.Model(&models.ProductsPurchased{}).Select("product_id, sum(product_quantity) as total").Where("order_id IN (?)", OrderId).Group("product_id").Order("total desc").Scan(&TopSellingProducts).Limit(conf.Product.TopProductCount)
	//.
	c.JSON(http.StatusOK, gin.H{"top_products": TopSellingProducts})
}
