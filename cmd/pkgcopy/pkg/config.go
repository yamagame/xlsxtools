package pkg

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Pkgs []*Package
}

func ReadConfig(config string) (*Config, error) {
	t := Config{}
	err := yaml.Unmarshal([]byte(config), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
