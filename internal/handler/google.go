package handler

import (
	"daistant-core/configs"
	"net/http"

	"github.com/gin-gonic/gin"

	"daistant-core/internal/service"
)

type GoogleHandler struct {
	config  *configs.Config
	service *service.GoogleService
}

func NewGoogleHandler(config *configs.Config, service *service.GoogleService) *GoogleHandler {
	return &GoogleHandler{config: config, service: service}
}

func (h *GoogleHandler) AuthGoogleCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
