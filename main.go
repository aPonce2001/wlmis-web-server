package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aPonce2001/wlmis-web-server/routes"
)

func main() {
	router := gin.Default()
	routes.ConfigureRoutes(router)
	router.Run(":5000")
}
