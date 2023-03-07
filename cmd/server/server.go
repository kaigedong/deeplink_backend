package main

import (
	"github.com/spf13/viper"

	"deeplink_backend/pkg/db"
	"deeplink_backend/pkg/log"
	"deeplink_backend/pkg/params"
	"deeplink_backend/pkg/router"
)

func init() {
	log.InitLogrus()
	params.InitParams()
	db.InitRedis()
}

func main() {
	engine := router.ConfigRouter()
	engine.Run(viper.GetString("http_port"))
}
