package Core

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string
	Owner    owner
	Database database
	Server   server
}

type database struct {
	DB      string
	Host    string
	Port    int
	Enabled bool
}

type server struct {
	Port int
}

type owner struct {
	Name string
}

func GetConfig(params ...string) Config {
	exists := false
	defaultConf := NewConfig()

	if len(params) > 0 {
		exists = ConfigFileExist(params[0])
	}

	if exists {
		conf := getConfigFromFile(params[0])
		return mergeConfig(defaultConf, conf)
	}

	return defaultConf

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
	defaultConf.Title = currentConf.Title
	defaultConf.Database = currentConf.Database
	defaultConf.Owner = currentConf.Owner
	defaultConf.Server = currentConf.Server

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
