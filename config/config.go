package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Database databaseConfig
	Server   serverConfig
	Redis    redisConfig
}

type serverConfig struct {
	Addr          string
	SessionSecret string `json:"session_secret"`
	SessionKey    string `json:"session_key"`
	Salt          string `json:"salt"`
}

type databaseConfig struct {
	DriverName string `json:"driver_name"`
	Host       string
	Port       string
	User       string
	Password   string
	Database   string
	DSN        string
}

type redisConfig struct {
	Host     string
	Port     string
	Password string
	MaxIdle  int `json:"max_idle"`
	Addr     string
}

var (
	Config   config
	Server   *serverConfig
	Database *databaseConfig
	Redis    *redisConfig
)

func initJson() {
	var configName string = "config"
	if os.Getenv("MODE") != "release" {
		configName += "_debug"
	}
	data, err := ioutil.ReadFile(fmt.Sprintf("config/%s.json", configName))
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &Config); err != nil {
		log.Fatal(err)
	}
}
func initServer() {
	Server = &Config.Server
}

func initDatabase() {
	Database = &Config.Database
	switch Database.DriverName {
	case "mysql":
		Database.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			Database.User,
			Database.Password,
			Database.Host,
			Database.Port,
			Database.Database,
		)
	}
}

func initRedis() {
	Redis = &Config.Redis
	Redis.Addr = Redis.Host + ":" + Redis.Port
}

func init() {
	initJson()
	initServer()
	initDatabase()
	initRedis()
}
