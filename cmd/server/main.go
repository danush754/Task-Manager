package main

import (
	"fmt"

	handler "github.com/danush754/Task-Manager/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	//  set up the server

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "pong")
	})

	h := handler.NewHandler()

	RegisterRoutes(r, h)

	r.Run(":8000")
}
