package config

import (
	"encoding/json"
	"github.com/redsofa/middleware-test/logger"
	"os"
)

var ServerConf Config

type Config struct {
	Port     int
	AuthPass bool
}

func ReadServerConf() {
	f, err := os.Open("server_config.json")
	defer f.Close()

	if err != nil {
		logger.Error.Println(err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&ServerConf)
	if err != nil {
		logger.Error.Println(err)
	}
}
