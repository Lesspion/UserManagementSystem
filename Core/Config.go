package Core

import (
	"log"
	"os"
	"strconv"

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
	Port    uint8
	Enabled bool
}

type server struct {
	Port uint8
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
	db := database{DB: getEnv("UMS_DATBASE_DB", "mongodb"), Host: getEnv("UMS_DATBASE_HOST", "localhost"), Port: getUint8Env("UMS_DATABASE_PORT", "27017"), Enabled: getBoolEnv("UMS_DATBASE_ENABLED", "true")}

	Server := server{Port: getUint8Env("UMS_SERVER_PORT", "8080")}

	Owner := owner{Name: getEnv("UMS_OWNER_NAME", "pawndev")}

	return Config{
		Title:    getEnv("UMS_TITLE", "UMS"),
		Database: db,
		Server:   Server,
		Owner:    Owner,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getBoolEnv(key, defaultValue string) bool {
	env, err := strconv.ParseBool(getEnv(key, defaultValue))
	if err != nil {
		log.Fatal(err)
	}

	return env
}

func getUint8Env(key, defaultValue string) uint8 {
	env, err := strconv.ParseUint(getEnv(key, defaultValue), 10, 8)
	if err != nil {
		log.Fatal(err)
	}

	return uint8(env)
}
