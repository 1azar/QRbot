package main

import (
	"QRbot/internal/qrbot"
	"flag"
	"log"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("Reading token")
	token := flag.String("token", "TELEGRAM_APITOKEN_QRBOT", "Bot token or the name of a system variable containing the token")
	tmp := os.Getenv(*token)
	if tmp != "" {
		token = &tmp
	}

	qrbot.StartBot(*token, infoLog, errorLog)
}
