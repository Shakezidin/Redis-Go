package main

import (
	"github.com/Shakezidin/configs"
	"github.com/Shakezidin/routes"
	"github.com/gin-gonic/gin"
)

var env = configs.Loadenv()
var R = gin.Default()

func init() {
	configs.GetRedis()
}
func main() {

	routes.Routes(R)

	R.Run(env.Port)

}
