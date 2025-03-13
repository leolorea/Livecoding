package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const token = "Bearer ds%IOF2e2!D&@gd#dsa#hulwG(*d(@98d29`*d@Y*)"

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		if tokenString != token {
			c.JSON(http.StatusForbidden, gin.H{"error": "Wrong token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func uploadFiles(c *gin.Context) {
	var meta MetaData

	err := c.ShouldBind(&meta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	// Get the image file from the form
	file, header, err := c.Request.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file data received"})
		log.Println(err)
		return
	}
	defer file.Close()

	// Read the binary data from the file
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file data"})

		return
	}

	meta.Data = data
	meta.Name = header.Filename
	InsertFile(meta)

}

// listFiles lists all uploaded images
func listFiles(c *gin.Context) {

	files := getFiles()

	c.JSON(http.StatusOK, files)

}
