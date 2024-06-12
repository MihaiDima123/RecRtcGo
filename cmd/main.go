package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/onchunk", func(c *gin.Context) {
		file, err := os.OpenFile("record.webm", os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			fmt.Print("Cannot open/create the file")
			return
		}
		defer file.Close()

		chunk, _, err := c.Request.FormFile("chunk")
		if err != nil {
			fmt.Print("Cannot get the chunk")
			return
		}
		defer chunk.Close()

		if _, err := io.Copy(file, chunk); err != nil {
			fmt.Print("Failed to write chunk to file")
			return
		}
	})

	err := r.Run()
	if err != nil {
		return
	}
}
