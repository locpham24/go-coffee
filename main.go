package main

import (
	"fmt"
	"github.com/locpham24/go-coffee/app/handler"
	"github.com/locpham24/go-coffee/config"
)

func main() {
	engine := handler.InitEngine()

	configs, err := config.LoadConfig("config")
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	fmt.Println("db username", configs.PostgreSQL.Username)
	engine.Run()
}
