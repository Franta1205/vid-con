package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"vid-con/routes"
)

func main() {
	router := gin.Default()
	fmt.Println("is working")

	routes.SetRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
