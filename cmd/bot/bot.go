package main

import (
	"log"

	"github.com/joho/godotenv"
	botPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command/add"
	cmdDeletePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command/delete"
	cmdHelpPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command/update"
	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
)

func runBot() {
	var book bookPkg.Interface
	{
		book = bookPkg.New()
	}
	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()
		commandAdd := cmdAddPkg.New(book)
		bot.RegisterHandler(commandAdd)
		commandList := cmdListPkg.New(book)
		commandUpdate := cmdUpdatePkg.New(book)
		bot.RegisterHandler(commandUpdate)
		bot.RegisterHandler(commandList)
		commandDelete := cmdDeletePkg.New(book)
		bot.RegisterHandler(commandDelete)
		commandHelp := cmdHelpPkg.New(map[string]string{
			commandAdd.Name():    commandAdd.Description(),
			commandList.Name():   commandList.Description(),
			commandDelete.Name(): commandDelete.Description(),
			commandUpdate.Name(): commandUpdate.Description(),
		})
		bot.RegisterHandler(commandHelp)
	}
	if err := bot.Run(); err != nil {
		log.Panic(err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}
	runBot()
}
