package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func CheckConfig() error {
	if viper.GetString("api_key") != "" {
		return nil
	}
	return SetConfig()
}

func SetConfig() error {

	fmt.Println("Setting up configuration")

	var apiKey string
	fmt.Print("Enter API Key: ")
	fmt.Scanln(&apiKey)
	if apiKey == "" {
		return errors.New("API Key cannot be empty")
	}
	viper.Set("api_key", apiKey)
	viper.Set("model_name", "gemini-pro")

	return nil

}
