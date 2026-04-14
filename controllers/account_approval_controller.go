package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func AddApproval(c *gin.Context) {
	var accountApproval models.User
	if err := c.ShouldBindJSON(&accountApproval); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "custom message": "failed to bind the json \n in the add approval function"})
		return
	}
	if err := services.AddApproval(accountApproval.Status,accountApproval.Email,accountApproval.Role); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User account status updated ", "status": accountApproval.Status})

}
