package main

import (
	"log"

	"captioner.com.ng/config"
	"captioner.com.ng/internal/captioner/adapters/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	config := app.App{
		Cfg: cfg,
	}
	app.New(config)
}
