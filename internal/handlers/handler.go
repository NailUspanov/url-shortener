package handlers

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.create)
		api.GET("/:short_url", h.find)
	}
	return router
}
