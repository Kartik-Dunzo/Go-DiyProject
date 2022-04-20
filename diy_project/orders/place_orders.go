package orders

import (
	"diy_project/config"
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var conf *config.Config

func BuyProducts(c *gin.Context) {
	// Validate input
	conf = &config.Config_parse
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.CartProducts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No products in cart !!!"})
		return
	}

	if len(input.CartProducts) > conf.Product.MaxProductsCart {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't place order more than " + strconv.Itoa(conf.Product.MaxProductQuantitySold) + " products in the cart  !!!"})
		return
	}
	//validating products in cart
	var Product models.Products

	for _, value := range input.CartProducts {

		if value.ProductQuantity > conf.Product.MaxProductQuantitySold {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Products cannot have quantity more than 5 ar one time !!!!"})
			return
		}

		if err := service.DbConn.Where("id = ?", value.ProductID).First(&Product).Error; err != nil {
			//errors := fmt.Sprintf("Product with id %s not found!", string(value.ProductID))
			c.JSON(http.StatusBadRequest, gin.H{"error": "product with id " + strconv.Itoa(value.ProductID) + "not found !!!!"})
			return
		}
		if (value.ProductQuantity + Product.Sold_Quantity) > Product.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product with id " + strconv.Itoa(value.ProductID) + "do not have quantity available for purchase !!!!"})
			return
		}

	}

	for _, value := range input.CartProducts {
		service.DbConn.Where("id = ?", value.ProductID).First(&Product)

		Product.Sold_Quantity = Product.Sold_Quantity + value.ProductQuantity
		if Product.Sold_Quantity == Product.Quantity {
			Product.Status = "Out_Of_Stock"
		}
		service.DbConn.Model(&Product).Update(Product)
	}
	input.CreatedAt = time.Now()
	service.DbConn.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": "Order created successfully. with Id : " + strconv.Itoa(input.Id)})

}
