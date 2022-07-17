package main

import (
	"log"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/commander"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/handlers"
)

func main() {
	log.Println("start main")
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}
	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
