package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Lock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}

	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("Name")
	defaultConfig.Username = os.Getenv("Username")
	defaultConfig.Password = os.Getenv("Password")
	defaultConfig.Address = os.Getenv("Address")
	cnv, err = strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}

	defaultConfig.Port = cnv

	return &defaultConfig
}
