package user

import (
	"diy_project/models"
	"diy_project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUser(c *gin.Context) {
	// Validate input
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.Users{Name: input.Name, Email: input.Email, Type: input.Type, Status: "active"}
	service.DbConn.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
