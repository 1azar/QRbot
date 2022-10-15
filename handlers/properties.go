// // Package handlers contains an entity that provides fields and handlers which will e used by bot
// // languages, logger instances, DB model, etc.
package handlers

import "log"

// Properties is an main entity in the package.
// Contain fields to be used by handler functions of the bot.
type Properties struct {
	Token   *string
	Lang    map[string]BotsWords
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

type BotsWords struct {
	Greeting string   `json:"greeting" validate:"required"`
	ReadyMsg []string `json:"ready_msg"  validate:"required"`
}

//// TODO несколько языков реализовать. пока тестовая реализация:
//func (p Properties) New(bw BotsWords) Properties {
//
//}
