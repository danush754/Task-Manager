package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//  set up the server

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "pong")
	})

	r.Run(":8000")
}
