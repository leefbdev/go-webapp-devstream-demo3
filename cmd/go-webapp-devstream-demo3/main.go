package main

import (
	"github.com/gin-gonic/gin"

	"github.com/leefbdev/go-webapp-devstream-demo3/internal/pkg/album"
)

func main() {
	router := gin.Default()

	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}
