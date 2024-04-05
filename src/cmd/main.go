package main

import (
	"log"

	"github.com/rezacloner1372/golang-clean-web-api/api"
	"github.com/rezacloner1372/golang-clean-web-api/config"
	"github.com/rezacloner1372/golang-clean-web-api/data/cache"
	"github.com/rezacloner1372/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)

}
