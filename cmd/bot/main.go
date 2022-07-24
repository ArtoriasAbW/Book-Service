package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/commander"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}
	log.Println("Bot starts")
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}
	handlers.AddHandlers(cmd)
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
