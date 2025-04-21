package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"parte3/api"
)

func main() {
	r := gin.Default()
	api.InitRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(fmt.Errorf("error trying to start server: %v", err))
	}
}
