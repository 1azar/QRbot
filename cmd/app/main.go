package main

import (
	"github.com/1azar/QRChan/infrastructure"
	"github.com/1azar/QRChan/interfaces/tgbot"
	"github.com/1azar/QRChan/usecases"
	"github.com/joho/godotenv"
	"os"
	"sync"
	"time"
)

func init() {
	if err := godotenv.Load("DEV.env"); err != nil {
		panic("No DEV.env file found") //fixme
	}
}

func main() {
	tokenVarName := "TELEGRAM_APITOKEN_QRBOT"

	qrInteractor := usecases.NewQRInteract(
		infrastructure.NewMemoryDBRepo(),
		infrastructure.NewLogger(),
		time.Duration(10),
	)

	TOKEN := os.Getenv(tokenVarName) // tg bot toke
	LANGDIR := "."                   // language directory where bot words wil be saved

	var wg sync.WaitGroup

	wg.Add(1)
	go tgbot.BotInitialize(TOKEN, LANGDIR, qrInteractor)

	wg.Wait()
}
