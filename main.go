package main

import (
	"github.com/akmamun/gin-boilerplate-examples/pkg/config"
	"github.com/akmamun/gin-boilerplate-examples/pkg/database"
	"github.com/akmamun/gin-boilerplate-examples/pkg/logger"
	"github.com/akmamun/gin-boilerplate-examples/routers"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.Connection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	db := database.GetDB()
	router := routers.Routes(db)

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
