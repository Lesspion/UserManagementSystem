package Core

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	title    string
	owner    owner
	database database
	server   server
}

type database struct {
	db      string
	host    string
	port    int
	enabled bool
}

type server struct {
	port int
}

type owner struct {
	name string
}

func GetConfig(params ...string) Config {
	exists := false
	defaultConf := NewConfig()

	if len(params) > 0 {
		exists = ConfigFileExist(params[0])
	}

	if exists {
		conf := getConfigFromFile(params[0])
		fmt.Print("billy")
		return mergeConfig(defaultConf, conf)
	} else {
		return defaultConf
	}
}

func ConfigFileExist(configfile string) bool {
	_, err := os.Stat(configfile)

	if err != nil {
		return false
	}

	return true
}

func getConfigFromFile(fileName string) Config {
	var config Config

	if _, err := toml.DecodeFile(fileName, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func mergeConfig(defaultConf Config, currentConf Config) Config {
	defaultConf.title = currentConf.title
	defaultConf.database = currentConf.database
	defaultConf.owner = currentConf.owner
	defaultConf.server = currentConf.server

	return defaultConf
}

func NewConfig() Config {
	return Config{
	// title: os.Getenv("UMS_TITLE"),
	// db:    os.Getenv("UMS_DB"),
	// host:  os.Getenv("UMS_DB_HOST"),
	// port:  os.Getenv("UMS_DB_PORT"),
	}
}
