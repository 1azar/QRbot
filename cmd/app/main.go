package main

import (
	"QRbot/handlers"
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	//"path/filepath"
)

func init() {
	if err := godotenv.Load("DEV.env"); err != nil {
		log.Println("No DEV.env file found") //fixme check api token for bot initializing
	}
}

func main() {
	// flags determination and parsing:
	token := flag.String("token", "TELEGRAM_APITOKEN_QRBOT", "Bot token or the name of a system variable containing the token")
	mongoURI := flag.String("mongo", "MONGODB_URI_QRBOT", "Mongo atlas uri or name of system variable containing connection uri")
	langDir := flag.String("langDir", ".", "Folder where localization jsons files are located.\n"+
		"filenames must be alphabetic language codes (example: ru.json, en.json).\n"+
		"To generate localization json template use -makeLangJSON=true flag")
	makeLangJSON := flag.Bool("makeLangJSON", false, "Generates in current directory localization JSON files named as en.json")
	debug := flag.Bool("debug", false, "Debug logs will be output to Os.stdout containing detailed information")

	flag.Parse()

	// props type contain all required settings and handlers for the bot.
	// upcoming steps below will set values for every field in the props type.
	props := handlers.Properties{}

	// initialise props loggers:
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	props.InfoLog = infoLog
	props.ErrLog = errorLog
	if *debug {
		debugLog := log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime)
		props.DebugEnabled = true
		props.DebLog = debugLog
		props.InfoLog.Println("Debug mode is active")
	}

	// generate bot responses dictionary if '-makeLangJSON' are provided
	if *makeLangJSON { //todo temp solution here, refractoring is required
		generateJSONen()
		generateJSONru()
		currentDir, err := os.Getwd()
		if err != nil {
			panic("couldn't get dir path")
		}
		props.InfoLog.Println(" 'en.json' and 'ru.json' file has been generated: ", currentDir)
		//return
	}

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

	// reading mongoDB uri from system variable or flag value:
	props.InfoLog.Println("Reading MongoDB connection URI..")
	tmp = os.Getenv(*mongoURI)
	if tmp != "" {
		mongoURI = &tmp
	}
	if *mongoURI == "" {
		props.ErrLog.Fatal("Can not read mongoDB URI.")
	}
	// initializing data base
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*mongoURI))
	if err != nil {
		props.ErrLog.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			props.ErrLog.Fatal(err)
		}
	}()
	props.DB.Collection = client.Database("QRBot").Collection("QRSettings")
	// Ping the primary:
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		props.ErrLog.Fatal(err)
	}
	props.InfoLog.Println("Data Base Successfully connected and pinged.")
	// test inserting
	rez, err := props.DB.CreateDefaultSettings(1)
	if err != nil {
		props.ErrLog.Fatal(err)
	}
	fmt.Println(rez)

	//searching and reading localization .json files:
	readLocalizationJson(&props, langDir)
	if len(props.Lang) == 0 {
		props.ErrLog.Fatal("Couldn't find any valid localisation json for the bot. " +
			"Pleas see -help for '-makeLangJSON' flag.")
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
