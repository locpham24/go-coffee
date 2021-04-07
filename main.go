package main

import (
	"github.com/locpham24/go-coffee/app/handler"
	"github.com/locpham24/go-coffee/config"
	"github.com/locpham24/go-coffee/infra"
)

func main() {
	engine := handler.InitEngine()

	configs, err := config.LoadConfig("config")
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	infra.InitPostgreSQL(configs)

	infra.InitLogging()

	engine.Run()
}
