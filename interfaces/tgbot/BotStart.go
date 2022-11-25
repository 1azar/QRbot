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
	btnQRTypeNormal := props.OptionsSelector.Data("⬜Normal", "btnQRTypeNormal")
	btnQRTypeWithLogo := props.OptionsSelector.Data("🍀With Logo", "btnQRTypeWithLogo")
	btnQRTypeHalftone := props.OptionsSelector.Data("🌓Halftone", "btnQRTypeHalftone")

	// Cell Shape submenu buttons:
	btnCellShapeCircle := props.OptionsSelector.Data("🏐Circle", "btnCellShapeCircle")
	btnCellShapeSquare := props.OptionsSelector.Data("🥅Square", "optCellShapeRectangle")

	// No choose submitting submenu: fixme temporary
	btnYesMaam := props.OptionsSelector.Data("Yes ma'am ಥ_ಥ", "optYesMaam")

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
		props.QRTypeSelector.Row(btnQRTypeNormal),
		props.QRTypeSelector.Row(btnQRTypeWithLogo),
		props.QRTypeSelector.Row(btnQRTypeHalftone),
	)
	// Choosing Cell Shape layout
	props.CellShapeSelector.Inline(
		props.CellShapeSelector.Row(btnCellShapeCircle),
		props.CellShapeSelector.Row(btnCellShapeSquare),
	)
	// No Choose layout
	props.NoChooseSelector.Inline(
		props.NoChooseSelector.Row(btnYesMaam),
	)

	/*
		// Choosing Color layout
		props.ColorSelector.Inline(
			props.ColorSelector.Row(optBGColorBlack, optBGColorWhite, optBGColorRed),
			props.ColorSelector.Row(optBGColorGreen, optBGColorBlue, optBGColorPurple),
			props.ColorSelector.Row(optBGColorOrange, optBGColorYellow, optBGColorBrown),
		)
	*/
	//	b.Handle(tele.OnText, func(c tele.Context) error {
	//	return c.Send("Hello!")
	//})

	// commands:
	b.Handle("/start", props.OnstartChat)
	b.Handle("/options", props.StartOptionsMenu)

	// inline buttons:
	// Enter option menu:
	b.Handle(&btnGetOptions, props.StartOptionsMenu)
	//// OPTIONS->QR TYPE:
	b.Handle(&btnQRType, props.ChangeQRType)
	b.Handle(&btnQRTypeNormal, props.ChangeQRTypeToNormal)
	b.Handle(&btnQRTypeWithLogo, props.ChangeQRTypeToWithLogo) // TODO unfinished
	b.Handle(&btnQRTypeHalftone, props.ChangeQRTypeToHalftone) // TODO unfinished
	// OPTIONS->CELL SHAPE:
	b.Handle(&btnQRCellShape, props.ChangeCellShape)
	b.Handle(&btnCellShapeCircle, props.ChangeCellShapeToCircle)
	b.Handle(&btnCellShapeSquare, props.ChangeCellShapeToSquare)
	// OPTIONS->BG COLOR
	b.Handle(&btnQRBGColor, props.ChangeBGColor) // TODO unfinished
	// OPTIONS->FG COLOR
	b.Handle(&btnQRFGColor, props.ChangeFGColor) // TODO unfinished
	//// OPTIONS->BORDER SIZE
	//b.Handle(&btnQRBGColor, props.ChangeBorderSize)
	// OPTIONS -> EXIT
	b.Handle(&btnExitOptions, props.ExitOptions)
	// OPTIONS -> Yes Ma'am (returns to option menu)
	b.Handle(&btnYesMaam, props.NoChoose)

	b.Handle(tele.OnText, props.OnText)

	props.BotInfoLog.Println("Bot started")
	b.Start()
}
