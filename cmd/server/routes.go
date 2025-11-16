package main

import (
	handler "github.com/danush754/Task-Manager/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {

	r.GET("/health", h.HealthCheck)

	r.POST("/sign-up", h.SignUp)

	r.POST("/login", h.Login)

	t := r.Group("/task")

	t.GET("/list", h.GetTaskList)

	t.POST("/create-task", h.CreateTask)

	t.GET("/list/:id", h.GetListDetails)

	t.PUT("/update-task/:id", h.UpdateTask)

}
