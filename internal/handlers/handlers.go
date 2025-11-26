package handler

import (
	"github.com/danush754/Task-Manager/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {

	return &Handler{Queries: queries}
}

func (h *Handler) HealthCheck(c *gin.Context) {

}

func (h *Handler) GetTaskList(c *gin.Context) {

}

func (h *Handler) CreateTask(c *gin.Context) {

}

func (h *Handler) GetListDetails(c *gin.Context) {

}

func (h *Handler) UpdateTask(c *gin.Context) {

}
