package main

import "go-calculate-service/router"

func main() {
	r := router.Setup()
	r.Run(":8081")
}
