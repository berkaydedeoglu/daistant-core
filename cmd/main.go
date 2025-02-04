package main

import (
	"daistant-core/configs"
	"daistant-core/internal/database"
	"daistant-core/internal/handler"
	"daistant-core/internal/repository"
	"daistant-core/internal/routing"
	"daistant-core/internal/service"
	"daistant-core/pkg/googleClient"
)

func main() {
	config := configs.New()

	db := database.New(config)
	database.Migrate(db)

	repo := repository.NewRepository(db)
	googleClient := googleClient.NewGoogleClient(config.GoogleOAuth.ClientID, config.GoogleOAuth.ClientSecret, config.GoogleOAuth.RedirectURL)
	service := service.NewGoogleService(config, repo, googleClient)
	handler := handler.NewGoogleHandler(config, service)
	router := routing.New(config, handler)
	router.RegisterThirdPartyRoutes()

	router.Run()
}
