package handlers

type Properties struct {
	lang Localisations
}

type Localisations struct {
	ru BotsWords
}

type BotsWords struct {
	Greeting string   `json:"greeting"`
	ReadyMsg []string `json:"ready_msg"`
}

// TODO несколько языков реализовать. пока тестовая реализация:
func (p Properties) New(bw BotsWords) Properties {

}
