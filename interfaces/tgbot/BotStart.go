package tgbot

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

func StartBot(props *Properties) {
	props.BotInfoLog.Println("Bot initializing")
	pref := tele.Settings{
		Token:  props.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		props.BotErrLog.Fatal(err)
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
	btnExitOptions := props.OptionsSelector.Data("✅SAVE", "btnExitOptions")

	// QR TYPE submenu buttons:
	optQRTypeNormal := props.OptionsSelector.Data("⬜Normal", "optQRTypeNormal")
	optQRTypeWithLogo := props.OptionsSelector.Data("🍀With Logo", "optQRTypeWithLogo")
	optQRTypeHalftone := props.OptionsSelector.Data("🌓Halftone", "optQRTypeHalftone")

	// Cell Shape submenu buttons:
	optCellShapeCircle := props.OptionsSelector.Data("🏐Circle", "optCellShapeCircle")
	optCellShapeSquare := props.OptionsSelector.Data("🥅Square", "optCellShapeRectangle")

	/*
		// Color submenu buttons:
		optBGColorBlack := props.OptionsSelector.Data("🖤Black", "optBGColorBlack")
		optBGColorWhite := props.OptionsSelector.Data("🤍White", "optBGColorWhite")
		optBGColorRed := props.OptionsSelector.Data("❤Red", "optBGColorRed")
		optBGColorGreen := props.OptionsSelector.Data("💚Green", "optBGColorGreen")
		optBGColorBlue := props.OptionsSelector.Data("💙Blue", "optBGColorBlue")
		optBGColorPurple := props.OptionsSelector.Data("💜Purple", "optBGColorPurple")
		optBGColorOrange := props.OptionsSelector.Data("🧡Orange", "optBGColorOrange")
		optBGColorYellow := props.OptionsSelector.Data("💛Yellow", "optBGColorYellow")
		optBGColorBrown := props.OptionsSelector.Data("🤎Brown", "optBGColorBrown")
	*/

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
	// Choosing Cell Shape layout
	props.CellShapeSelector.Inline(
		props.CellShapeSelector.Row(optCellShapeCircle),
		props.CellShapeSelector.Row(optCellShapeSquare),
	)

	/*
		// Choosing Color layout
		props.ColorSelector.Inline(
			props.ColorSelector.Row(optBGColorBlack, optBGColorWhite, optBGColorRed),
			props.ColorSelector.Row(optBGColorGreen, optBGColorBlue, optBGColorPurple),
			props.ColorSelector.Row(optBGColorOrange, optBGColorYellow, optBGColorBrown),
		)
	*/

	// commands:
	b.Handle("/start", props.OnstartChat)
	b.Handle("/options", props.StartOptionsMenu)

	// inline buttons:
	// Enter option menu:
	b.Handle(&btnGetOptions, props.StartOptionsMenu)
	//// OPTIONS->QR TYPE:
	b.Handle(&btnQRType, props.ChangeQRType)
	b.Handle(&optQRTypeNormal, props.ChangeQRTypeToNormal)
	b.Handle(&optQRTypeWithLogo, props.ChangeQRTypeToWithLogo)
	//b.Handle(&optQRTypeHalftone, props.ChangeQRTypeToHalftone)
	// OPTIONS->CELL SHAPE:
	b.Handle(&btnQRCellShape, props.ChangeCellShape)
	b.Handle(&optCellShapeCircle, props.ChangeCellShapeToCircle)
	b.Handle(&optCellShapeSquare, props.ChangeCellShapeToSquare)
	//// OPTIONS->BG COLOR
	//b.Handle(&btnQRBGColor, props.ChangeBGColor)
	//// OPTIONS->FG COLOR
	//b.Handle(&btnQRBGColor, props.ChangeFGColor)
	//// OPTIONS->BORDER SIZE
	//b.Handle(&btnQRBGColor, props.ChangeBorderSize)
	// OPTIONS -> EXIT
	b.Handle(&btnExitOptions, props.ExitOptions)

	//b.Handle(tele.OnText, MessageProcessing)

	props.BotInfoLog.Println("Bot started")
	b.Start()
}
