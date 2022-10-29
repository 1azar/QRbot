/*
	This file contains only functions which handles
	all but commands or buttons.
	Those functions must be methods of Properties type
	to gain axes to required fields (like languages etc.)
*/

package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

// OnstartChat method process /start command
func (p Properties) OnstartChat(c tele.Context) error {
	p.debugLog(fmt.Sprintf("OnstartChat for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	//TODO show examples of qr as pics (add property to p containing photos)
	return c.Send(p.BotResponses(c.Sender().LanguageCode).Greeting, &p.GeneralSelector)
}

// StartOptionsMenu method process click on Options button
// and respond with options menu in terms of inline buttons
func (p Properties) StartOptionsMenu(c tele.Context) error {
	p.debugLog(fmt.Sprintf("StartOptionsMenu for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
}

// ChangeQRType method process click on QR Type option button
// and respond with list of inline buttons which represent all possible types for qr
func (p Properties) ChangeQRType(c tele.Context) error {
	p.debugLog(fmt.Sprintf("ChangeQRType for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	return c.Edit(p.BotResponses(c.Sender().LanguageCode).QRTypeMsg, &p.QRTypeSelector)
}

// ChangeQRTypeToNormal method set "normal" qr type
func (p Properties) ChangeQRTypeToNormal(c tele.Context) error {
	p.debugLog(fmt.Sprintf("ChangeQRTypeToNormal for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	//TODO change User QR Type to "normal"
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
}

// ChangeQRTypeToWithLogo method set "with logo" qr type
func (p Properties) ChangeQRTypeToWithLogo(c tele.Context) error {
	p.debugLog(fmt.Sprintf("ChangeQRTypeToWithLogo for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	//TODO change User QR Type to "with logo"
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
}

// ChangeQRTypeToHalftone method set "halftone" qr type
func (p Properties) ChangeQRTypeToHalftone(c tele.Context) error {
	p.debugLog(fmt.Sprintf("ChangeQRTypeToHalftone for user:%X\t[%s]", c.Sender().ID, c.Sender().LanguageCode))
	//TODO change User QR Type to "halftone"
	return c.Send(p.BotResponses(c.Sender().LanguageCode).OptionsMsg, &p.OptionsSelector)
}
