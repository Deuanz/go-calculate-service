package calculator

import (
	"go-calculate-service/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sub(c *gin.Context) {
	data := entities.Data{}
	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := data.A.Sub(data.B)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
