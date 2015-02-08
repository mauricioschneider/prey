package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host      string
	Protocol  string
	ApiKey    string
	DeviceKey string
	Plugins   map[string]*Plugin
}

type Plugin struct {
	Name string
	opts map[string]string
}

func main() {
	config := getConfig()

}

func getConfig() Config {
	config := Config{}
	config.Host = getConfigValue("host")
	config.ApiKey = getConfigValue("api_key")
	config.DeviceKey = getConfigValue("device_key")
	config.Protocol = getConfigValue("protocol")
	return config
}
func getConfigValue(key string) string {
	viper.SetConfigName("prey")
	viper.ReadInConfig()
	return viper.GetStringMapString("control-panel")[key]
}
