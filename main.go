// main.go
package main

import (
	"tron-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/balance/:address", handlers.GetBalance)
	router.GET("/transactions/:address", handlers.GetTransactions)
	router.Run(":8080")
}
