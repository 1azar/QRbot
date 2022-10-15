package main

import (
	"QRbot/handlers"
	"QRbot/internal/qrbot"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-playground/validator"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"strings"
	//"path/filepath"
)

func main() {
	//TODO decompose main:

	// flags determination and parsing:
	token := flag.String("token", "TELEGRAM_APITOKEN_QRBOT", "Bot token or the name of a system variable containing the token")
	langDir := flag.String("langDir", ".", "Folder where localization jsons files are located.\n"+
		"filenames must be alphabetic language codes (example: ru.json, en.json).\n"+
		"To generate localization json template use -makeLangJSON=true flag")
	makeLanguageJSON := flag.Bool("makeLangJSON", false, "Generates in current directory localization JSON files named as en.json")
	flag.Parse()

	fmt.Println(langDir, makeLanguageJSON)

	// props type contain all required settings and handlers for the bot.
	// upcoming steps below will set values for every field in the props type.
	props := handlers.Properties{}

	// initialise loggers:
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
		props.InfoLog.Println(fmt.Sprintf("Checking for %s file \n", file.Name()))
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

	qrbot.StartBot(*token, infoLog, errorLog)
}

// // findFile searches for files whose names match  any language code and have the "ext" extension
func findFile(targetDir string, ext string) {

}
