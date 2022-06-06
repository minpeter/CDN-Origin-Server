package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "file upload failed",
			"detail":  err.Error(),
		})
		return
	}
	save_file, err := os.Create("./bucket/" + "new.png")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "file upload failed",
			"detail":  err.Error(),
		})
		return
	}

	defer save_file.Close()

	io.Copy(save_file, file)

	c.JSON(200, gin.H{
		"message": "file upload success",
	})
}
