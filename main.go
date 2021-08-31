package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application")

	router := gin.Default()

	router.GET("/albums", getAllAlbum)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", addAlbum)

	router.Run("localhost:8080")
}
