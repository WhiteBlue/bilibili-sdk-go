package main

import (
	"github.com/whiteblue/bilibili-go/service"
	"github.com/go-playground/log"
)

func main() {
	app, err := service.NewApplication("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	app.Router.Run(":8080")
}
