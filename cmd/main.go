package main

import (
	"log"

	"github.com/begenov/register-service/internal/app"
	"github.com/begenov/register-service/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error cfg: %v", err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal("error application\t", err)
	}

}
