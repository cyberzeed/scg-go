package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func setupConfig(appName string) *viper.Viper {
	var configFile string
	flag.StringVar(&configFile, "config", defaultConfigFile, "Configuration file")
	flag.Parse()

	config := viper.New()
	config.SetConfigName(appName)
	config.SetConfigType("yaml")
	config.SetConfigFile(configFile)
	config.AddConfigPath(fmt.Sprintf("/etc/%v", appName))
	config.AddConfigPath(fmt.Sprintf("$HOME/.%v", appName))
	config.AddConfigPath(".")
	config.SetDefault("server.address", defaultServerAddress)
	config.SetDefault("server.port", defaultServergPort)
	config.SetDefault("google.apiKey", defaultGoogleAPIKey)
	config.SetDefault("linebot.accessToken", defaultLineBotAccessToken)
	config.SetDefault("linebot.secret", defaultLineBotSecret)
	config.WriteConfig()
	config.BindPFlags(flag.CommandLine)
	return config
}
