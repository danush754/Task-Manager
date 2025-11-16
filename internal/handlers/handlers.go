package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {

	return &Handler{}
}

func (h *Handler) HealthCheck(c *gin.Context) {

}

func (h *Handler) SignUp(c *gin.Context) {

}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) GetTaskList(c *gin.Context) {

}

func (h *Handler) CreateTask(c *gin.Context) {

}

func (h *Handler) GetListDetails(c *gin.Context) {

}

func (h *Handler) UpdateTask(c *gin.Context) {

}
