package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	Mongo struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Username string `yaml:"username"`
	} `yaml:"mongo"`

	Redis struct {
		Addr string `yaml:"addr"`
	} `yaml:"redis"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
