package main

import (
	"cadanapay/controllers"
	"fmt"
)

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	server := gin.Default()

	// server routes
	server.POST("/api/v1/rates/conversion", controllers.GetExchangeRate)

	err := server.Run(":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
}
