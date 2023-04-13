package handler

import (
	"booking-service/internal/services/sendservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	sendService sendservice.ISendService
}

func NewHandler(sendService sendservice.ISendService) *Handler {
	return &Handler{
		sendService: sendService,
	}
}

func (h *Handler) ConfigAPIRoute(router *gin.Engine) {
	routers := router.Group("v1")
	routers.GET("employee", h.getAvailableEmployee())
}
