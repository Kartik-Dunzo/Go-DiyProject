package product

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductsById(c *gin.Context) {
	// Check Product exist
	var product models.Products
	if err := service.DbConn.Where("id = ?", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
