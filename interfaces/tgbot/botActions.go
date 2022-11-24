package tgbot

import (
	"fmt"
	"github.com/1azar/QRChan/domain"
	tele "gopkg.in/telebot.v3"
)

// processQRSettingsSearch firstly search for qr settings in the buffer and returns it if succeeded
// Secondly it searches in the repository and returns qr setting if succeeded
// Lastly, if there is no qr setting nether in buffer or repository, then it generates default qr settings and returns it
func (p Properties) processQRSettingsSearch(id int64) (domain.QRSettings, error) {
	var err error
	err = nil
	qs, ok := p.QRInteractor.QRSettingsBuffer.Get(id)
	if !ok { // if there is no QRSetting fou current user in buffer
		p.BotInfoLog.Println(fmt.Sprintf("Could not find QR Settings in buffer for user: %v", id))
		qs, err = p.QRInteractor.FindQRSettings(id) // search for qr setting in the db (if there is no qr settings for user then it will generate default settings)
	}
	return qs, err
}

// handlerWrapper firstly process searching qr settings for current user (from cache, db or generates new).
// secondly invokes passed function f (f functions changes qr settings parameters for current user)
// lastly saves qr settings of current user to cache and generates text representation of qr settings
func (p Properties) handlerWrapper(f func(settings *domain.QRSettings)) func(id int64) (msg string, err error) {
	return func(id int64) (msg string, err error) {
		qs, err := p.processQRSettingsSearch(id)
		f(&qs)
		p.QRInteractor.QRSettingsBuffer.Set(qs.ID, qs)
		msg = generateOptionsMenuText(qs)
		return
	}
}

// OnstartChat method process /start command
func (p Properties) OnstartChat(c tele.Context) error {
	//TODO show examples of qr as pics (add property to p containing photos)
	p.BotInfoLog.Println(fmt.Sprintf("/start for user: %v", c.Sender().ID))
	return c.Send(p.BotResponses(c.Sender().LanguageCode).Greeting, &p.GeneralSelector)
}

// StartOptionsMenu method process click on Options button
// and respond with options menu in terms of inline buttons
func (p Properties) StartOptionsMenu(c tele.Context) error {
	qs, err := p.processQRSettingsSearch(c.Sender().ID)
	if err != nil {
		return err
	}
	p.QRInteractor.QRSettingsBuffer.Set(qs.ID, qs) // now qr setting is buffered for current user.
	msg := generateOptionsMenuText(qs)
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ExitOptions method process click on Exit button
// and sent to repos QR Settings for current context
func (p Properties) ExitOptions(c tele.Context) error {
	qs, err := p.processQRSettingsSearch(c.Sender().ID)
	if err != nil {
		return err
	}
	if err := p.QRInteractor.StoreQRSettings(qs); err != nil {
		return err
	}
	msg := generateOptionsMenuText(qs)
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).GeneralMsg+msg, &p.GeneralSelector, tele.ModeHTML)
}

// ChangeQRType method process click on QR Type option button
// and respond with list of inline buttons which represent all possible types for qr
func (p Properties) ChangeQRType(c tele.Context) error {
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).QRTypeMsg, &p.QRTypeSelector)
}

// ChangeQRTypeToNormal method set "normal" qr type
func (p Properties) ChangeQRTypeToNormal(c tele.Context) error {
	msg, err := p.handlerWrapper(func(qs *domain.QRSettings) {
		qs.QRType.Name = domain.Normal
	})(c.Sender().ID)
	if err != nil {
		return err
	}
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ChangeQRTypeToWithLogo method set "with logo" qr type
func (p Properties) ChangeQRTypeToWithLogo(c tele.Context) error {
	//TODO change User QR Type to "with logo". not finished yet has to request for logo image
	msg, err := p.handlerWrapper(func(qs *domain.QRSettings) {
		qs.QRType.Name = domain.WithLogo
	})(c.Sender().ID)
	if err != nil {
		return err
	}
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ChangeQRTypeToHalftone method set "halftone" qr type
func (p Properties) ChangeQRTypeToHalftone(c tele.Context) error {
	//TODO change User QR Type to "halftone". not finished yet has to request for image
	msg, err := p.handlerWrapper(func(qs *domain.QRSettings) {
		qs.QRType.Name = domain.Halftone
	})(c.Sender().ID)
	if err != nil {
		return err
	}
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ChangeCellShape method process click on Cell Shape option button
// and respond with list of inline buttons which represent all possible cell shapes for QR
func (p Properties) ChangeCellShape(c tele.Context) error {
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).CellShapeMsg, &p.CellShapeSelector)
}

// ChangeCellShapeToCircle method set "Circle" qr shape
func (p Properties) ChangeCellShapeToCircle(c tele.Context) error {
	msg, err := p.handlerWrapper(func(qs *domain.QRSettings) {
		qs.CellShape = domain.Circle
	})(c.Sender().ID)
	if err != nil {
		return err
	}
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ChangeCellShapeToSquare method set "Square" qr shape
func (p Properties) ChangeCellShapeToSquare(c tele.Context) error {
	msg, err := p.handlerWrapper(func(qs *domain.QRSettings) {
		qs.CellShape = domain.Rectangle
	})(c.Sender().ID)
	if err != nil {
		return err
	}
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg+"\n"+msg, &p.OptionsSelector, tele.ModeHTML)
}

// ChangeBGColor method process click on BG Color option button
func (p Properties) ChangeBGColor(c tele.Context) error {
	//return c.Edit(p.BotResponses(c.Sender().LanguageCode).BGColorMsg, &p.ColorSelector)
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).BGColorMsg)
}

//// ChangeBGColorTo method process qr bg color changing
//func (p Properties) ChangeBGColorTo(c tele.Context) error {
//	//TODO change User cell shape to "circle"
//	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
//}

// ChangeFGColor method process click on FG Color option button
func (p Properties) ChangeFGColor(c tele.Context) error {
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).BGColorMsg)
}

//// ChangeFGColorTo method process qr fg color changing
//func (p Properties) ChangeFGColorTo(c tele.Context) error {
//	//TODO change User cell shape to "circle"
//	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
//}
