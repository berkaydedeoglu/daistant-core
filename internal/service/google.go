package service

import (
	"daistant-core/configs"
	"daistant-core/internal/model/http/response"
	"daistant-core/internal/repository"

	"github.com/gin-gonic/gin"
)

type GoogleService interface {
	GetOAuthURL(ctx *gin.Context, userId uint) *response.GoogleOAuthURLResponse
}

type googleService struct {
	config *configs.Config
	repo   *repository.Repository
}

func NewGoogleService(config *configs.Config, repo *repository.Repository) GoogleService {
	return &googleService{config: config, repo: repo}
}

func (s *googleService) GetOAuthURL(ctx *gin.Context, userId uint) *response.GoogleOAuthURLResponse {
	resp := response.GoogleOAuthURLResponse{
		BaseURL:      s.config.GoogleOAuth.OAuthURL,
		Scope:        s.config.GoogleOAuth.Scope,
		ClientID:     s.config.GoogleOAuth.ClientID,
		RedirectURI:  s.config.GoogleOAuth.RedirectURL,
		ResponseType: "code",
	}

	return resp.SetOAuthURL(userId)
}
