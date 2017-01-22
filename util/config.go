package util

import (
	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	User string `toml:"db"`
	Pass string `toml:"pass"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func GetDBConfig() (*DBConfig, error) {
	var conf DBConfig
	if _, err := toml.DecodeFile("../_tools/config.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
