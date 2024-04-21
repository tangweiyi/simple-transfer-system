package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.AddConfigPath("..")     // call multiple times to add many search paths
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func ReadConfig(path string, result any) error {
	v := viper.Sub(path)
	fmt.Println(viper.AllKeys())
	return v.Unmarshal(result)
}
