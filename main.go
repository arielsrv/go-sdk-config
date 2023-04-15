package main

import (
	"log"

	"github.com/go-sdk-config/config"
)

func main() {
	config.NewBuilder().
		WithFile("config.yml").
		WithFolder("resources").
		Build()

	log.Println(config.String("base.url"))
	log.Println(config.String("app.env"))
	log.Println(config.String("host"))
	log.Println(config.String("public"))
}
