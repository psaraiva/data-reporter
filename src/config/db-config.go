package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DBConfig struct {
	Engine   string
	Host     string
	Port     string
	Database string
	User     string
	Pass     string
}

func (c *DBConfig) Load(dbengine string) (DBConfig, error) {
	config := DBConfig{}
	dbfile := fmt.Sprintf("./config/database.%s.json", dbengine)
	if _, err := os.Stat(dbfile); err != nil {
		return config, err
	}

	data, err := os.ReadFile(dbfile)
	if err != nil {
		return config, err
	}

	if json.Unmarshal(data, &config); err != nil {
		log.Println(err)
	}

	return config, nil
}
