package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type data struct {
	A decimal.Decimal `json:"a"`
	B decimal.Decimal `json:"b"`
}

func sum(c *gin.Context) {
	data := data{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := data.A.Add(data.B)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func sub(c *gin.Context) {
	data := data{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := data.A.Sub(data.B)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func mul(c *gin.Context) {
	data := data{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := data.A.Mul(data.B)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func div(c *gin.Context) {
	data := data{}
	if err := c.ShouldBindJSON(&data); err != nil {
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

func main() {
	r := gin.Default()
	r.POST("/calculator.sum", sum)
	r.POST("/calculator.sub", sub)
	r.POST("/calculator.mul", mul)
	r.POST("/calculator.div", div)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
