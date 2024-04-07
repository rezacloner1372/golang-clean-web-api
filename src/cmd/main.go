package main

import (
	"github.com/rezacloner1372/golang-clean-web-api/api"
	"github.com/rezacloner1372/golang-clean-web-api/config"
	"github.com/rezacloner1372/golang-clean-web-api/data/cache"
	"github.com/rezacloner1372/golang-clean-web-api/data/db"
	"github.com/rezacloner1372/golang-clean-web-api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)

}
