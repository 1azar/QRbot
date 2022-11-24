package tgbot

import (
	"encoding/json"
	"fmt"
	"github.com/1azar/QRChan/domain"
	"github.com/go-playground/validator"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readLocalizationJson(props *Properties, langDir string) {
	props.BotInfoLog.Println("Searching for localisation JSON files..")
	files, err := os.ReadDir(langDir)
	if err != nil {
		props.BotErrLog.Fatal(err)
	}
	var bw BotsWords
	validate := validator.New()
	tmpLang := make(map[string]BotsWords)
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
		props.BotInfoLog.Println(fmt.Sprintf("Checking for %s file", file.Name()))
		byteJson, err := os.ReadFile(filepath.Join(langDir, file.Name()))
		if err != nil {
			props.BotErrLog.Println(err)
			continue
		}
		err = json.Unmarshal(byteJson, &bw)
		if err != nil {
			props.BotErrLog.Println(err)
			continue
		}
		err = validate.Struct(bw)
		if err != nil {
			props.BotInfoLog.Println(fmt.Sprintf("%s file has invalid structure. Can not be used", file.Name()))
			continue
		}
		tmpLang[langID] = bw
	}
	props.Lang = tmpLang

}

// Создание ответов бота:
func generateJSONen() {
	RuBotResponses := BotsWords{
		Greeting: "Hi! I am QR-chan ☜(ﾟヮﾟ☜).\n" +
			"You can use me to generate QR codes." +
			"[can do nothing but this 👀‍]\n" +
			"You send me message - I will encode it  t̶h̶e̶ ̶h̶e̶c̶k̶ ̶u̶p̶ gracefully.\n" +
			"Also i have some options. You will figure out how to use them." +
			"You are smart human [maybe].",
		GeneralMsg: "Ready to generate! QR Parameters:\n",
		ReadyMsg: []string{"Here you go:",
			"Get this, human:",
			"Take a QRevich's square:",
			"I hope I didn't mix up any bits:",
			"Why is no one paying me for this:",
			"I forgot the 2nd law of robotics ;-;\n I hope there is nothing related to QR, otherwise I will be fired.",
			"I hope I didn't code anything evil:",
			"01101001 00100000 01101100 01101111 01110110 01100101 00100000 01111001 01101111 01110101:"},
		OptionsMsg:   "Set me up, set me up completely",
		QRTypeMsg:    "Choose what type of QR to create",
		CellShapeMsg: "Choose QR Cell shape",
		BGColorMsg:   "Choose QR background color",
		FGColorMsg:   "Choose QR face ground color",
	}
	file, _ := json.MarshalIndent(RuBotResponses, "", " ")
	_ = ioutil.WriteFile("en.json", file, 0644)
}

// Создание ответов бота:
func generateJSONru() {
	RuBotResponses := BotsWords{
		Greeting: "Привет! Меня зовут QR-чан [с пловом ☜(ﾟヮﾟ☜)].\n" +
			"Можешь пользоваться мной для генерации QR кодов." +
			"[я больше ничего не умею 🤷‍]\n" +
			"Пришлешь мне сообщения - я его закодирую н̶а̶х̶у̶й̶ бережно.\n" +
			"Еще у меня настройки есть, как ими пользоваться, думаю, разберешься." +
			"Ты же умный человек [наверное].",
		GeneralMsg: "Готова генерровать! Праметры QR:\n",
		ReadyMsg: []string{"Вот, готово:",
			"Держи, кожанный:",
			"Забирай квадрат QRевича:",
			"Надеюсь никакой бит не перепутала:",
			"И почему мне за это никто не платит:",
			"Забыла 2ой закон робототехники ;-;\n Надеюсь там нет ничего связанного с QR, инче меня уволят.",
			"Надеюсь я ничего плохого не закодировала:",
			"01101001 00100000 01101100 01101111 01110110 01100101 00100000 01111001 01101111 01110101:"},
		OptionsMsg:   "Настривай, настривай меня полностью",
		QRTypeMsg:    "Выбери какой тип QR мне создавать",
		CellShapeMsg: "Выбери какой тип ячеек для QR",
		BGColorMsg:   "Выбери цвет фона QR",
		FGColorMsg:   "Выбери цвет QR",
	}
	file, _ := json.MarshalIndent(RuBotResponses, "", " ")
	_ = ioutil.WriteFile("ru.json", file, 0644)
}

func generateOptionsMenuText(qs domain.QRSettings) string {
	txt := [][]string{
		[]string{"QR Type", qs.QRType.Name.String()},
		[]string{"Cell Shape", qs.CellShape.String()},
		[]string{"BG Color", fmt.Sprintf("(%v, %v, %v)", qs.BackGroundColor.R, qs.BackGroundColor.G, qs.BackGroundColor.B)},
		[]string{"FG Color", fmt.Sprintf("(%v, %v, %v)", qs.ForeGroundColor.R, qs.ForeGroundColor.G, qs.ForeGroundColor.B)},
		[]string{"Border Size", fmt.Sprintf("%v %s", qs.BorderWidth.Value, qs.BorderWidth.WidthType)},
	}

	txtBuf := new(strings.Builder)
	table := tablewriter.NewWriter(txtBuf)

	for _, v := range txt {
		table.Append(v)
	}
	table.Render()
	msg := "<pre>" + txtBuf.String() + "</pre>"
	return msg

	//var txt string
	//txt += fmt.Sprintf("%17s | %12s\n", "QR Type", qs.QRType.Name)
	//txt += fmt.Sprintf("%17s | %12s\n", "Cell Shape", qs.CellShape)
	//txt += fmt.Sprintf("%17s | %12s\n", "BG Color", fmt.Sprintf("(%v, %v, %v) \n", qs.BackGroundColor.R, qs.BackGroundColor.G, qs.BackGroundColor.B))
	//txt += fmt.Sprintf("%17s | %12s\n", "FG Color", fmt.Sprintf("(%v, %v, %v) \n", qs.ForeGroundColor.R, qs.ForeGroundColor.G, qs.ForeGroundColor.B))
	//txt += fmt.Sprintf("%17s | %12s\n", "Border Size", fmt.Sprintf("%v %s \n", qs.BorderWidth.Value, qs.BorderWidth.WidthType))
	//return txt
}
