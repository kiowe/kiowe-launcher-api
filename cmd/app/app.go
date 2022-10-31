package main

import (
	"github.com/kiowe/kiowe-launcher-api/internal/app"
	"log"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	a.Run()
}
