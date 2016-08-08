package main

import (
	"github.com/go-playground/log"
	"github.com/whiteblue/bilibili-go/service"
)

func main() {
	app, err := service.NewApplication("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	app.Router.Run(":8080")
}
