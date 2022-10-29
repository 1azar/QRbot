package qrbot

import (
	"QRbot/handlers"
	"time"

	tele "gopkg.in/telebot.v3"
)

//var (
//	selector   = &tele.ReplyMarkup{}
//	btnOptions = selector.Data("🐱Options", "btnOptions")
//)

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

	// buttons construction //todo add other languages for buttons name
	// button to get into options inline menu:
	btnGetOptions := props.GeneralSelector.Data("⚙Options", "btnGetOptions")

	// option menu buttons:
	btnQRType := props.OptionsSelector.Data("🕸QR Type", "btnQRType")
	btnQRCellShape := props.OptionsSelector.Data("⏺Cell shape", "btnQRCellShape")
	btnQRBGColor := props.OptionsSelector.Data("🌈BG color", "btnQRBGColor")
	btnQRFGColor := props.OptionsSelector.Data("🌈FG color", "btnQRFGColor")
	btnQRBorderSizer := props.OptionsSelector.Data("🔳Border size", "btnQRBorderSizer")
	btnExitOptions := props.OptionsSelector.Data("✅Exit", "btnExitOptions")

	optQRTypeNormal := props.OptionsSelector.Data("⬜Normal", "optQRTypeNormal")
	optQRTypeWithLogo := props.OptionsSelector.Data("🍀With Logo", "optQRTypeWithLogo")
	optQRTypeHalftone := props.OptionsSelector.Data("🌓Halftone", "optQRTypeHalftone")

	// layouts for GetOption button
	props.GeneralSelector.Inline(
		props.GeneralSelector.Row(btnGetOptions),
	)
	// layouts for Options buttons
	props.OptionsSelector.Inline(
		props.GeneralSelector.Row(btnQRType, btnQRCellShape),
		props.GeneralSelector.Row(btnQRBGColor, btnQRFGColor),
		props.GeneralSelector.Row(btnQRBorderSizer, btnExitOptions),
	)
	// Choosing QR type layout
	props.QRTypeSelector.Inline(
		props.QRTypeSelector.Row(optQRTypeNormal),
		props.QRTypeSelector.Row(optQRTypeWithLogo),
		props.QRTypeSelector.Row(optQRTypeHalftone),
	)

	// commands:
	b.Handle("/start", props.OnstartChat)
	b.Handle("/options", props.StartOptionsMenu)

	// inline buttons
	b.Handle(&btnGetOptions, props.StartOptionsMenu)

	b.Handle(&btnQRType, props.ChangeQRType)
	b.Handle(&optQRTypeNormal, props.ChangeQRTypeToNormal)
	b.Handle(&optQRTypeWithLogo, props.ChangeQRTypeToWithLogo)
	b.Handle(&optQRTypeHalftone, props.ChangeQRTypeToHalftone)

	//b.Handle(tele.OnText, MessageProcessing)

	props.InfoLog.Println("Bot started")
	b.Start()
}
