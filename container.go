package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func GetuserGin(c *gin.Context) {
	c.JSON(http.StatusOK, "Hey, this is Pradeep! Have Wonderfull day!!!")
}

func main() {
	fmt.Print("Hello World!")
	router.GET("/aboutme", GetuserGin)
	router.Run(":8080")
}
