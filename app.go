package main

import (
	"fmt"
	"log"

	"github.com/fvbock/endless"
	flag "github.com/spf13/pflag"
)

const (
	defaultConfigFile         = "scg.yaml"
	defaultGoogleAPIKey       = ""
	defaultLineBotAccessToken = ""
	defaultLineBotSecret      = ""
	defaultServerAddress      = "0.0.0.0"
	defaultServergPort        = 8080
)

func init() {
	flag.String("google.apikey", "", "Google API key")
	flag.String("linebot.accesstoken", "", "Access token of Line Bot")
	flag.String("linebot.secret", "", "Secret of Line Bot")
	flag.String("server.address", defaultServerAddress, "Listening address")
	flag.Int("server.port", defaultServergPort, "Listening port")
}

func main() {
	const appName = "scg"

	log.Println("Load configuration")
	config := setupConfig(appName)
	address := config.GetString("server.address")
	port := config.GetInt("server.port")
	listen := fmt.Sprintf("%v:%v\n", address, port)
	router := setupRouter(config)

	log.Println(fmt.Sprintf("Listen %v\n", listen))
	router.Run()
	endless.ListenAndServe(listen, router)
}
