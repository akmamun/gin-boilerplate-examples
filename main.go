package main

import (
	"github.com/akmamun/gin-boilerplate-examples/pkg/config"
	"github.com/akmamun/gin-boilerplate-examples/pkg/database"
	"github.com/akmamun/gin-boilerplate-examples/pkg/logger"
	"github.com/akmamun/gin-boilerplate-examples/pkg/routers"
	"github.com/spf13/viper"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config.SetupConfig() error: %s", err)
	}

	if err := database.Connection(); err != nil {
		logger.Fatalf("database.DbConnection error: %s", err)
	}

	db := database.GetDB()
	router := routers.Routes(db)

	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8000")

	logger.Fatalf("%v:%v", router.Run(viper.GetString("server.host")+":"+viper.GetString("server.port")))

}
