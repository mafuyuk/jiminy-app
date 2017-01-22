package util

import (
	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func GetDBConfig() (*DBConfig, error) {
	var conf DBConfig
	if _, err := toml.Decode("_tools/config.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
