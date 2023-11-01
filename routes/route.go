package routes

import (
	"github.com/Shakezidin/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	c.GET("/", controllers.Home)
	c.POST("/signup", controllers.SignupPost)
	c.POST("/login", controllers.LoginPost)
}
