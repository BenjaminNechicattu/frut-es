package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	mysqlservice "frutes/database/mysql"
	"frutes/utils"
)

type ServerConfig struct {
	Port        *int                     `json:"port" validate:"required"`
	LogFilePath *string                  `json:"logfilepath"`
	MysqlConfig mysqlservice.MysqlConfig `json:"mysqlconfig" validate:"required"`
}

func ParseServerConfig(configFilePath string) (*ServerConfig, error) {
	log.Print(configFilePath)

	rawjsonbytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var severConfig ServerConfig

	err = json.Unmarshal(rawjsonbytes, &severConfig)
	if err != nil {
		log.Print("description : ", err)
		return nil, err
	}

	msg, err := utils.ValidateStruct(severConfig)
	if err != nil {
		log.Print("validation failed : ", msg)
		return nil, err
	}

	return &severConfig, nil
}

func GetServerConfig() (*ServerConfig, error) {
	var configFileName string
	flag.StringVar(&configFileName, "configfile", "", "Please pass the config file")

	flag.Parse()
	if configFileName == "" {
		log.Fatal("Please pass the config file as the argument")
	}
	log.Println("Config File:", configFileName)
	return ParseServerConfig(configFileName)
}
