package main

import (
	"image/png"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func images(c *gin.Context) {

	height, _ := strconv.Atoi(c.Query("h"))
	width, _ := strconv.Atoi(c.Query("w"))
	imageName := c.Param("image")

	file, err := os.Open("./bucket/" + imageName)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "image open failed",
			"detail":  err.Error(),
		})
		return
	}

	imageData, err := png.Decode(file)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "image decode failed",
		})
		return
	}

	file, err = os.Create("./resize/" + imageName)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "image resize failed",
			"detail":  err.Error(),
		})
		return
	}
	png.Encode(file, resize.Resize(uint(width), uint(height), imageData, resize.NearestNeighbor))
	c.File("./resize/" + imageName)
	defer file.Close()
}
