package qrbot

import (
	"QRbot/handlers"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot(props *handlers.Properties) {
	props.InfoLog.Println("Bot initializing")
	pref := tele.Settings{
		Token:  *props.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		props.ErrLog.Fatal(err)
		return
	}

	//props.InfoLog.Println("Reading bot response dictionary")

	//b.Handle("/start", OnstartChat)

	b.Handle(tele.OnText, MessageProcessing)

	b.Start()
}
