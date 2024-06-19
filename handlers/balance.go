package handlers

import (
	"net/http"
	"tron-api/tron"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	address := c.Param("address")
	client := tron.NewTronClient()
	balance, err := client.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
