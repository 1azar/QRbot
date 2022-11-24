package tgbot

import (
	"fmt"
	"github.com/1azar/QRChan/usecases"
	"log"
	"os"
)

// BotInitialize initializes the tg bot.
// token - system variable name containing token or token itself
// langDir - directory where to save and load bots bot response dictionary
func BotInitialize(token, langDir string, qrInteractor *usecases.QRInteract) {
	// flags determination and parsing:
	//token := flag.String("token", "TELEGRAM_APITOKEN_QRBOT", "Bot token or the name of a system variable containing the token")
	//mongoURI := flag.String("mongo", "MONGODB_URI_QRBOT", "Mongo atlas uri or name of system variable containing connection uri")
	//langDir := flag.String("langDir", ".", "Folder where localization jsons files are located.\n"+
	//	"filenames must be alphabetic language codes (example: ru.json, en.json).\n"+
	//	"To generate localization json template use -makeLangJSON=true flag")
	//makeLangJSON := flag.Bool("makeLangJSON", false, "Generates in current directory localization JSON files named as en.json")
	//debug := flag.Bool("debug", false, "Debug logs will be output to Os.stdout containing detailed information")

	//flag.Parse()

	// props type contain all required settings and handlers for the bot.
	// upcoming steps below will set values for every field in the props type.
	props := Properties{}

	// initialise props loggers:
	infoLog := log.New(os.Stdout, "[BOT] INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[BOT] ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	props.BotInfoLog = infoLog
	props.BotErrLog = errorLog

	// generate bot responses dictionary if '-makeLangJSON' are provided
	// now bot always generates response dictionary. TODO implement some choosing mechanism
	if true {
		generateJSONen()
		generateJSONru()
		currentDir, err := os.Getwd()
		if err != nil {
			panic("couldn't get dir path")
		}
		props.BotInfoLog.Println(" 'en.json' and 'ru.json' file has been generated: ", currentDir)
	}

	// reading token from system variable or flag value:
	props.BotInfoLog.Println("Reading token..")
	tmpT := os.Getenv(token)
	if tmpT != "" {
		token = tmpT
	}
	if token == "" {
		props.BotErrLog.Fatal("Can not use empty token.")
	}
	props.Token = token

	// connecting QRInteractor:
	props.BotInfoLog.Println("Connecting QR Interactor..")
	props.QRInteractor = qrInteractor

	//searching and reading localization .json files:
	readLocalizationJson(&props, langDir)
	if len(props.Lang) == 0 {
		props.BotErrLog.Fatal("Couldn't find any valid localisation json for the bot.")
	}
	props.BotInfoLog.Println(fmt.Sprintf("Bot will spek next languages: %s",
		func() string {
			var str string
			for k := range props.Lang {
				str += k + " "
			}
			return str
		}()))

	//fmt.Println("FINISHED")
	StartBot(&props)
}
