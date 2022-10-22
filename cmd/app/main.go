package main

import (
	"QRbot/handlers"
	"flag"
	"fmt"
	"log"
	"os"
	//"path/filepath"
)

func main() {
	// flags determination and parsing:
	token := flag.String("token", "TELEGRAM_APITOKEN_QRBOT", "Bot token or the name of a system variable containing the token")
	langDir := flag.String("langDir", ".", "Folder where localization jsons files are located.\n"+
		"filenames must be alphabetic language codes (example: ru.json, en.json).\n"+
		"To generate localization json template use -makeLangJSON=true flag")
	makeLanguageJSON := flag.Bool("makeLangJSON", false, "Generates in current directory localization JSON files named as en.json")
	flag.Parse()

	// generate bot responses dictionary if '-makeLangJSON' are provided
	if *makeLanguageJSON {
		generateJSONen()
		currentDir, err := os.Getwd()
		if err != nil {
			panic("couldn't get dir path")
		}
		fmt.Println(" 'en.json' file has been generated: ", currentDir)
		return
	}

	// props type contain all required settings and handlers for the bot.
	// upcoming steps below will set values for every field in the props type.
	props := handlers.Properties{}

	// initialise props loggers:
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	props.InfoLog = infoLog
	props.ErrLog = errorLog

	// reading token from system variable or flag value:
	props.InfoLog.Println("Reading token..")
	tmp := os.Getenv(*token)
	if tmp != "" {
		token = &tmp
	}
	if *token == "" {
		props.ErrLog.Fatal("Can not read token.")
	}
	props.Token = token

	//searching and reading localization .json files:
	readLocalizationJson(&props, langDir)
	if len(props.Lang) == 0 {
		props.ErrLog.Fatal("Couldn't find any valid localisation json for the bot. " +
			"Pleas see -help for '-makeLanguageJSON' flag.")
	}
	props.InfoLog.Println(fmt.Sprintf("Bot will spek next languages: %s",
		func() string {
			var str string
			for k := range props.Lang {
				str += k + " "
			}
			return str
		}()))

	fmt.Println("FINISHED")
	//qrbot.StartBot(&props)
}
