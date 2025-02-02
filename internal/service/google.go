package service

import (
	"daistant-core/configs"
	"daistant-core/internal/repository"
)

type GoogleService struct {
	config *configs.Config
	repo   *repository.Repository
}

func NewGoogleService(config *configs.Config, repo *repository.Repository) *GoogleService {
	return &GoogleService{config: config, repo: repo}
}
