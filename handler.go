package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAllAlbum responds with the list of all albums as JSON.
func getAllAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// addAlbum add new album from received JSON request.
func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbum locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
