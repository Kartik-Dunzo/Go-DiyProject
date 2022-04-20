package product

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateProducts(c *gin.Context) {
	// Validate product exist
	var product models.Products
	if err := service.DbConn.Where("id = ?", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	// Validate input
	var input models.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Updating the data of product
	service.DbConn.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
