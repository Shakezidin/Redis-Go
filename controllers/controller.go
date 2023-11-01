package controllers

import (
	"github.com/Shakezidin/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

type Signup struct {
	Email    string
	Password string
}

func SignupPost(c *gin.Context) {
	var user Signup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}
	if err := models.RegisterUser(user.Email); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}
func LoginPost(c *gin.Context) {
	var user Signup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}
	if err := models.AuthenticateUser(user.Email); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status": true,
		"email":  user.Email,
	})
}
