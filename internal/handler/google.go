package handler

import (
	"daistant-core/configs"
	"daistant-core/internal/model/http/request"
	"net/http"

	"github.com/gin-gonic/gin"

	"daistant-core/internal/service"
)

type GoogleHandler struct {
	config  *configs.Config
	service service.GoogleService
}

func NewGoogleHandler(config *configs.Config, service service.GoogleService) *GoogleHandler {
	return &GoogleHandler{config: config, service: service}
}

func (h *GoogleHandler) AuthGoogleCallback(c *gin.Context) {
	var req request.GoogleAuthCallbackRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ExchangeCode(c, 1, req.Code, req.Scope)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (h *GoogleHandler) GetOAuthURL(c *gin.Context) {
	resp := h.service.GetOAuthURL(c, 1)
	c.JSON(http.StatusOK, resp)
}
