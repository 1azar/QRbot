// // Package handlers contains an entity that provides fields and handlers which will e used by bot
// // languages, logger instances, DB model, etc.
package handlers

import (
	"QRbot/models"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
)

// Properties is an main entity in the package.
// Contain fields to be used by handler functions of the bot.
type Properties struct {
	Token             *string
	DB                models.UserSettingModel
	Lang              map[string]BotsWords //Lang keys is language code (ex:"en","ru", etc.)
	InfoLog           *log.Logger
	ErrLog            *log.Logger
	DebugEnabled      bool
	DebLog            *log.Logger
	GeneralSelector   tele.ReplyMarkup //GeneralSelector contains inline buttons attached to each message from the bot
	OptionsSelector   tele.ReplyMarkup //OptionsSelector contains inline buttons that appear when entering the bot menu
	QRTypeSelector    tele.ReplyMarkup //QRTypeSelector contains inline buttons that represent list of possible types for qr
	CellShapeSelector tele.ReplyMarkup //CellShapeSelector contains inline buttons that represent list of possible cell shapes for qr
	ColorSelector     tele.ReplyMarkup //ColorSelector contains inline buttons that represent list of possible colors for qr
}

// BotResponses method returns pointer to BotsWords included in Lang field of Properties type.
func (p Properties) BotResponses(LanguageCode string) *BotsWords {
	// if bot can speak "LanguageCode" language
	if bw, ok := p.Lang[LanguageCode]; ok {
		p.debugLog(fmt.Sprintf("BotResponses for [%s]:\tok", LanguageCode))
		return &bw
	}
	// else if bot can speak "en" language
	if bw, ok := p.Lang["en"]; ok {
		p.debugLog(fmt.Sprintf("BotResponses for [%s]:\tnot ok, choosed [en]", LanguageCode))
		return &bw
	}
	// else if bot can speak any language
	for key, bw := range p.Lang {
		p.debugLog(fmt.Sprintf("BotResponses for [%s]:\tnot ok, couldnt find [end], choosed [%s]", LanguageCode, key))
		return &bw
	}
	// else if bot does not speak language
	// (should never happen because it is checked during initialization in main.go)
	return &BotsWords{}
}

func (p Properties) debugLog(msg string) {
	if p.DebugEnabled {
		p.DebLog.Println(msg)
	}
}

type BotsWords struct {
	Greeting     string   `json:"greeting" validate:"required"`       //ReadyMsg: bot say that when command /start is used
	ReadyMsg     []string `json:"ready_msg"  validate:"required"`     //ReadyMsg: bot say some of that when sending the result to user
	OptionsMsg   string   `json:"options_msg" validate:"required"`    //OptionsMsg: bot say that when user entering the bot options menu
	QRTypeMsg    string   `json:"qr_type_msg" validate:"required"`    //QRTypeMsg: bot say that when user entering the qr type changing menu
	CellShapeMsg string   `json:"cell_shape_msg" validate:"required"` //TODO CellShapeMsg: bot say that when user entering the cell shape changing menu
	BGColorMsg   string   `json:"bg_color_msg" validate:"required"`   //TODO BGColorMsg: bot say that when user entering the BG color changing menu
	FGColorMsg   string   `json:"fg_color_msg" validate:"required"`   //TODO FGColorMsg: bot say that when user entering the FG color changing menu
}
