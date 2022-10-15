package qrbot

import (
	tele "gopkg.in/telebot.v3"
)

//func (l Localisations) OnstartChat(c tele.Context) error {
//
//	return c.Send("On Start msg")
//}

func MessageProcessing(c tele.Context) error {
	return c.Send(c.Sender().LanguageCode)
}

//func Logger(str string) {
//	now := time.Now()
//	nowStr := now.Format("02-01-2006 15:04:05")
//	fmt.Printf(">>>\t%s\t%s", nowStr, str)
//}
