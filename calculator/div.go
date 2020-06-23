package calculator

import (
	"net/http"

	"go-calculate-service/entities"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func Div(c *gin.Context) {
	data := entities.Data{}
	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.B.Equal(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "denominator can not be zero"})
		return
	}

	result := data.A.Div(data.B)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
