package product

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddLIstOfProducts(c *gin.Context) {
	// Validate input
	var input models.ListOfProducts
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Product_list) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No products to add!!!"})
		return
	}
	var check_user models.Users
	if err := service.DbConn.Where("id = ? AND status = 'active' AND type = 'merchant'", input.User_Id).First(&check_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is unauthorized to add Products!!!"})
		return
	}
	// using for loop
	for _, value := range input.Product_list {
		NewProduct := models.Products{User_Id: input.User_Id, Name: value.Name, Category: value.Category, Price: value.Price, Quantity: value.Quantity, Sold_Quantity: 0, Status: "active"}
		service.DbConn.Create(&NewProduct)
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Products added successfully !!!"})
}
