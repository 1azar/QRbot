package qrbot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot(token string, infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Println("Bot initializing")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	infoLog.Println("Reading bot response dictionary")

	//b.Handle("/start", OnstartChat)

	b.Handle(tele.OnText, MessageProcessing)

	b.Start()
}
