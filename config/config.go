package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server string
	Database string
}

func (c *Config) Read()  {
	if _,err := toml.Decode("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}