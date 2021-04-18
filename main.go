package main

import (
	"github.com/locpham24/go-coffee/app/handler"
	"github.com/locpham24/go-coffee/app/orm"
	"github.com/locpham24/go-coffee/config"
	"github.com/locpham24/go-coffee/infra"
	"os"
)

func main() {
	defer close()

	engine := handler.InitEngine()

	configs, err := config.LoadConfig("config")
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	infra.InitPostgreSQL(configs)
	orm.InitOrmInstances()

	infra.InitLogging()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine.Run(":" + port)
}

func close() {
	infra.ClosePostgreSql()
}
