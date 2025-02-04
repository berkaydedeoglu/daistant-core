package service

import (
	"daistant-core/configs"
	"daistant-core/internal/model"
	"daistant-core/internal/model/http/response"
	"daistant-core/internal/repository"
	"daistant-core/pkg/googleClient"
	"time"

	"github.com/gin-gonic/gin"
)

type GoogleService interface {
	GetOAuthURL(ctx *gin.Context, userId uint) *response.GoogleOAuthURLResponse
	ExchangeCode(ctx *gin.Context, userId uint, code string, scope string) error
	GetAccessToken(ctx *gin.Context, userId uint) (string, error)
}

type googleService struct {
	config       *configs.Config
	repo         repository.Repository
	googleClient googleClient.GoogleClient
}

func NewGoogleService(config *configs.Config, repo repository.Repository, googleClient googleClient.GoogleClient) GoogleService {
	return &googleService{config: config, repo: repo, googleClient: googleClient}
}

func (s *googleService) GetOAuthURL(ctx *gin.Context, userId uint) *response.GoogleOAuthURLResponse {
	resp := response.GoogleOAuthURLResponse{
		BaseURL:      s.config.GoogleOAuth.OAuthURL,
		Scope:        s.config.GoogleOAuth.Scope,
		ClientID:     s.config.GoogleOAuth.ClientID,
		RedirectURI:  s.config.GoogleOAuth.RedirectURL,
		ResponseType: "code",
		AccessType:   "offline",
	}

	return resp.SetOAuthURL(userId)
}

func (s *googleService) ExchangeCode(ctx *gin.Context, userId uint, code string, scope string) error {
	exchangeCodeResponse, err := s.googleClient.ExchangeCode(code)
	if err != nil {
		return err
	}

	thirdPartyConnection := model.ThirdPartyConnection{
		Provider:     model.ProviderGoogle,
		AccessToken:  exchangeCodeResponse.AccessToken,
		RefreshToken: exchangeCodeResponse.RefreshToken,
		Scope:        scope,
		UserID:       userId,
		ExpiresAt:    time.Now().Add(time.Duration(exchangeCodeResponse.ExpiresIn) * time.Second),
	}

	return s.repo.CreateThirdPartyConnection(ctx, &thirdPartyConnection)
}

func (s *googleService) GetAccessToken(ctx *gin.Context, userId uint) (string, error) {
	tpc, err := s.repo.GetThirdPartyConnectionByUserID(ctx, userId, model.ProviderGoogle)
	if err != nil {
		return "", err
	}

	if tpc.ExpiresAt.Before(time.Now()) {
		err = s.refreshToken(ctx, userId, tpc.RefreshToken)
		if err != nil {
			return "", err
		}
	}

	return tpc.AccessToken, nil
}

func (s *googleService) refreshToken(ctx *gin.Context, userId uint, refreshToken string) error {
	exchangeCodeResponse, err := s.googleClient.RefreshToken(refreshToken)
	if err != nil {
		return err
	}

	thirdPartyConnection := model.ThirdPartyConnection{
		Provider:     model.ProviderGoogle,
		AccessToken:  exchangeCodeResponse.AccessToken,
		RefreshToken: exchangeCodeResponse.RefreshToken,
		UserID:       userId,
		ExpiresAt:    time.Now().Add(time.Duration(exchangeCodeResponse.ExpiresIn) * time.Second),
	}

	return s.repo.UpdateThirdPartyConnection(ctx, &thirdPartyConnection)
}
