package main

import (
	"github.com/akmamun/gin-boilerplate-examples/pkg/config"
	"github.com/akmamun/gin-boilerplate-examples/pkg/database"
	"github.com/akmamun/gin-boilerplate-examples/pkg/logger"
	"github.com/akmamun/gin-boilerplate-examples/routers"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, err := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	if err != nil {
		logger.Fatalf("set timezone of application: %s", err)
	}
	time.Local = loc

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
