package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shotenLinker/pkg/api/handler"
	"shotenLinker/pkg/config"
)

type ServerHTTP struct {
	engine *gin.Engine
	config config.Config
}

func NewServerHTTP(linkHandler *handler.LinksHandler, cfg config.Config) *ServerHTTP {
	engine := gin.New()

	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(gin.Logger())

	// Not really elegant decision
	engine.Use(func(c *gin.Context) {
		c.Set("AppPort", cfg.AppPort)
		c.Set("AppUrl", cfg.AppUrl)
		c.Next()
	})

	engine.POST("/encode", linkHandler.Encode)
	engine.GET("/decode/:shortURL", linkHandler.Decode)

	return &ServerHTTP{engine: engine, config: cfg}
}

func (sh *ServerHTTP) Start() {
	address := fmt.Sprintf(":%s", sh.config.AppPort)
	// Unhandled error
	sh.engine.Run(address)
}
