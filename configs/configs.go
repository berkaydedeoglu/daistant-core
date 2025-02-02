package configs

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func New() *Config {
	var k = koanf.New("../")

	err := k.Load(file.Provider("config.yaml"), yaml.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	out := Config{}
	k.Unmarshal("", &out)
	return &out
}
