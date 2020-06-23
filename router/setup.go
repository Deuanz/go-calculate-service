package router

import (
	"go-calculate-service/calculator"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.POST("/calculator.sum", calculator.Sum)
	r.POST("/calculator.sub", calculator.Sub)
	r.POST("/calculator.mul", calculator.Mul)
	r.POST("/calculator.div", calculator.Div)
	return r
}
