package product

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(c *gin.Context) {
	// Check Product exist
	var product []models.Products
	if err := service.DbConn.Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}
	if len(product) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Products available at the moment!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
