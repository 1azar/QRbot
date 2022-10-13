package handlers

import tele "gopkg.in/telebot.v3"

func () OnstartChat(c tele.Context) error {

	return c.Send("On Start msg")
}
