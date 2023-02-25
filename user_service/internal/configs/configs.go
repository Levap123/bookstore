package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Postgres struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Username string `yaml:"username"`
		DBName   string `yaml:"db_name"`
	} `yaml:"postgres"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
