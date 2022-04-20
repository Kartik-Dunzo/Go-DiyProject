package user

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	// Validate input
	var user models.Users
	if err := service.DbConn.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Updating the data of user
	service.DbConn.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
