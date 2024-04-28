package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"shotenLinker/pkg/api"
	"shotenLinker/pkg/api/handler"
	"shotenLinker/pkg/config"
	"shotenLinker/pkg/domain"
	"shotenLinker/pkg/repository"
	"shotenLinker/pkg/service"
)

func main() {
	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load cfg: ", configErr)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.Links{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	server := api.NewServerHTTP(
		handler.NewLinkHandler(
			service.NewLinkUseCase(
				repository.NewLinkRepository(db),
			),
		),
		cfg,
	)

	server.Start()
}
