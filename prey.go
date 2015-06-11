package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

const ClientVersion = "0.0.1"

type Config struct {
	Host      string
	Protocol  string
	ApiKey    string
	DeviceKey string
	Plugins   map[string]*Plugin
}

type Plugin struct {
	Name string
	Opts map[string]string
}

func main() {
	config := initConfig()
	commands, _ := getCommands(config)
	fmt.Println(commands)
}

func initConfig() Config {
	config := Config{}
	config.Host = readConfigValue("host")
	config.ApiKey = readConfigValue("api_key")
	config.DeviceKey = readConfigValue("device_key")
	config.Protocol = readConfigValue("protocol")
	return config
}

func readConfigValue(key string) string {
	viper.SetConfigName("prey")
	viper.ReadInConfig()
	return viper.GetStringMapString("control-panel")[key]
}

func getCommands(config Config) (string, error) {
	apiUrl := "/api/v2/devices/"
	apiFormat := ".json"
	commandsUrl := strings.Join([]string{config.Protocol, "://", config.Host, apiUrl, config.DeviceKey, apiFormat}, "")

	client := &http.Client{}
	req, err := http.NewRequest("GET", commandsUrl, nil)
	req.SetBasicAuth(config.ApiKey, "x")
	req.Header.Set("User-Agent", "Prey/"+ClientVersion)

	resp, err := client.Do(req)

	if err != nil {
		return "Unable to get commands", err
	} else {

		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "Unable to parse body", err
		}

		return string(contents), nil
	}
}
