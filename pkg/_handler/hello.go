package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	fmt.Println("hello")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, "Hello v1")
}
