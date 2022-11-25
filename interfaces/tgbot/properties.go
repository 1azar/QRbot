package tgbot

/*
	========================================================
	contains the structures necessary for the bot to work
	========================================================
*/
import (
	"github.com/1azar/QRChan/usecases"
	tele "gopkg.in/telebot.v3"
	"log"
)

// Properties is a main entity for the bot.
// Contain fields to be used by handler functions of the bot.
type Properties struct {
	Token             string
	QRInteractor      *usecases.QRInteract
	Lang              map[string]BotsWords //Lang keys is language code (ex:"en","ru", etc.)
	BotInfoLog        *log.Logger
	BotErrLog         *log.Logger
	GeneralSelector   tele.ReplyMarkup //GeneralSelector contains inline buttons attached to each message from the bot
	OptionsSelector   tele.ReplyMarkup //OptionsSelector contains inline buttons that appear when entering the bot menu
	QRTypeSelector    tele.ReplyMarkup //QRTypeSelector contains inline buttons that represent list of possible types for qr
	CellShapeSelector tele.ReplyMarkup //CellShapeSelector contains inline buttons that represent list of possible cell shapes for qr
	//ColorSelector     tele.ReplyMarkup //ColorSelector contains inline buttons that represent list of possible colors for qr
	NoChooseSelector tele.ReplyMarkup //NoChooseSelector contains inline button with only one agreement option
}

// BotResponses method returns pointer to BotsWords included in Lang field of Properties type.
func (p Properties) BotResponses(LanguageCode string) *BotsWords {
	// if bot can speak "LanguageCode" language
	if bw, ok := p.Lang[LanguageCode]; ok {
		return &bw
	}
	// else if bot can speak "en" language
	if bw, ok := p.Lang["en"]; ok {
		return &bw
	}
	// else if bot can speak any language
	for _, bw := range p.Lang {
		return &bw
	}
	// else if bot does not speak language
	// (should never happen because it is checked during initialization in main.go)
	return &BotsWords{}
}

type BotsWords struct {
	Greeting     string   `json:"greeting" validate:"required"`       //ReadyMsg: bot say that when command /start is used
	GeneralMsg   string   `json:"general_msg" validate:"required"`    //GeneralMsg: bot say that when it is ready to serve
	ReadyMsg     []string `json:"ready_msg"  validate:"required"`     //ReadyMsg: bot say some of that when sending the result to user
	OptionsMsg   string   `json:"options_msg" validate:"required"`    //OptionsMsg: bot say that when user entering the bot options menu
	QRTypeMsg    string   `json:"qr_type_msg" validate:"required"`    //QRTypeMsg: bot say that when user entering the qr type changing menu
	CellShapeMsg string   `json:"cell_shape_msg" validate:"required"` //TODO CellShapeMsg: bot say that when user entering the cell shape changing menu
	BGColorMsg   string   `json:"bg_color_msg" validate:"required"`   //TODO BGColorMsg: bot say that when user entering the BG color changing menu
	FGColorMsg   string   `json:"fg_color_msg" validate:"required"`   //TODO FGColorMsg: bot say that when user entering the FG color changing menu
}

// GetReadyMsg returns random element of ReadyMsg of BotsWords struct field
func (bw *BotsWords) GetReadyMsg() string {
	for _, bw := range bw.ReadyMsg {
		return bw
	}
	return bw.ReadyMsg[0]
}
