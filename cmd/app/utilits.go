// Contains the functions necessary to initialize the bot.
// Used in main.go file of the package
package main

import (
	"QRbot/handlers"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"golang.org/x/text/language"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readLocalizationJson(props *handlers.Properties, langDir *string) {
	props.InfoLog.Println("Searching for localisation JSON files..")
	files, err := os.ReadDir(*langDir)
	if err != nil {
		props.ErrLog.Fatal(err)
	}
	var bw handlers.BotsWords
	validate := validator.New()
	tmpLang := make(map[string]handlers.BotsWords)
	for _, file := range files {
		fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		fileExt := filepath.Ext(file.Name())
		langID := language.All.Make(fileName).String()
		if fileExt != ".json" {
			continue
		}
		if langID == "und" {
			continue
		}
		props.InfoLog.Println(fmt.Sprintf("Checking for %s file", file.Name()))
		byteJson, err := os.ReadFile(filepath.Join(*langDir, file.Name()))
		if err != nil {
			props.ErrLog.Println(err)
			continue
		}
		err = json.Unmarshal(byteJson, &bw)
		if err != nil {
			props.ErrLog.Println(err)
			continue
		}
		err = validate.Struct(bw)
		if err != nil {
			props.InfoLog.Println(fmt.Sprintf("%s file has invalid structure. Can not be used", file.Name()))
			continue
		}
		tmpLang[langID] = bw
	}
	props.Lang = tmpLang

}

// Создание ответов бота:
func generateJSONen() {
	RuBotResponses := handlers.BotsWords{
		Greeting: "Hi! I am QR-chan ☜(ﾟヮﾟ☜).\n" +
			"You can use me to generate QR codes." +
			"[can do nothing but this 👀‍]\n" +
			"You send me message - I will encode it  t̶h̶e̶ ̶h̶e̶c̶k̶ ̶u̶p̶ gracefully.\n" +
			"Also i have some options. You will figure out how to use them." +
			"You are smart human [maybe].",
		ReadyMsg: []string{"Here you go:",
			"Get this, human:",
			"Take a QRevich's square:",
			"I hope I didn't mix up any bits:",
			"Why is no one paying me for this:",
			"I forgot the 2nd law of robotics ;-;\n I hope there is nothing related to QR, otherwise I will be fired.",
			"I hope I didn't code anything evil:",
			"01101001 00100000 01101100 01101111 01110110 01100101 00100000 01111001 01101111 01110101:"},
		OptionsMsg: "Set me up, set me up completely",
		QRTypeMsg:  "Choose what type of QR to create",
	}
	file, _ := json.MarshalIndent(RuBotResponses, "", " ")
	_ = ioutil.WriteFile("en.json", file, 0644)
}

// Создание ответов бота:
func generateJSONru() {
	RuBotResponses := handlers.BotsWords{
		Greeting: "Привет! Меня зовут QR-чан [с пловом ☜(ﾟヮﾟ☜)].\n" +
			"Можешь пользоваться мной для генерации QR кодов." +
			"[я больше ничего не умею 🤷‍]\n" +
			"Пришлешь мне сообщения - я его закодирую н̶а̶х̶у̶й̶ бережно.\n" +
			"Еще у меня настройки есть, как ими пользоваться, думаю, разберешься." +
			"Ты же умный человек [наверное].",
		ReadyMsg: []string{"Вот, готово:",
			"Держи, кожанный:",
			"Забирай квадрат QRевича:",
			"Надеюсь никакой бит не перепутала:",
			"И почему мне за это никто не платит:",
			"Забыла 2ой закон робототехники ;-;\n Надеюсь там нет ничего связанного с QR, инче меня уволят.",
			"Надеюсь я ничего плохого не закодировала:",
			"01101001 00100000 01101100 01101111 01110110 01100101 00100000 01111001 01101111 01110101:"},
		OptionsMsg: "Настривай, настривай меня полностью",
		QRTypeMsg:  "Выбери какой тип QR мне создавать",
	}
	file, _ := json.MarshalIndent(RuBotResponses, "", " ")
	_ = ioutil.WriteFile("ru.json", file, 0644)
}
