package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

type Server struct {
	Port string `json:"port"`
}

type Database struct {
	ConnectionString string `json:"connectionString"`
}

func Load() Configs {
	var config Configs
	env := os.Getenv("env")
	configFile, err := os.Open("./configs/config." + env + ".json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	port := os.Getenv("PORT")

	if port != "" {
		config.Server.Port = port
	}
	return config
}
