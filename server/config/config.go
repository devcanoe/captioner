package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP  `yaml:"http"`
	MONGO `yaml:"mongo"`
	HASH  `yaml:"hash"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type HASH struct {
	PrivateKey string `env-required:"true" env:"JWT_PRIVATE_KEY" yaml:"private"`
	PublicKey  string `env-required:"true" env:"JWT_PUBLIC_KEY" yaml:"public"`
}

type MONGO struct {
	Uri      string `env-required:"true" yaml:"uri" env:"DATABASE_URL"`
	Database string `env-required:"true" yaml:"database" env:"DATABASE_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
