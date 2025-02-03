package routing

import (
	"daistant-core/configs"
	"daistant-core/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	config  *configs.Config
	router  *gin.Engine
	handler *handler.GoogleHandler
}

func New(config *configs.Config, handler *handler.GoogleHandler) *Router {
	gin.SetMode(config.GinMode)
	router := gin.Default()

	return &Router{
		config:  config,
		router:  router,
		handler: handler,
	}
}

func (r *Router) RegisterThirdPartyRoutes() {
	thirdParty := r.router.Group("/third-party")

	thirdParty.GET("/google/oauth/url", r.handler.GetOAuthURL)
	thirdParty.GET("/google/oauth/callback", r.handler.AuthGoogleCallback)
}

func (r *Router) Run() {
	r.router.Run(fmt.Sprintf(":%d", r.config.Port))
}
