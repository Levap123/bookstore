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

	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	JWTSign string `yaml:"jwt_sign"`

	Validator struct {
		PasswordMin int `yaml:"password_min"`
		PasswordMax int `yaml:"password_max"`

		UsernameMin int `yaml:"username_min"`
		UsernameMax int `yaml:"username_max"`
	} `yaml:"validator"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
