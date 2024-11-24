package config

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env              string     `yaml:"env"`
	Server           GrpcServer `yaml:"server"`
	ConnectionString string     `yaml:"connection_string"`
	TokenTTL         int        `yaml:"token_ttl"`
}

type GrpcServer struct {
	Address string        `yaml:"address"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Не возвращает ошибку, вызывается только при запуске приложения
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}

	return &cfg
}

func Load() (*Config, error) {
	path := fetchConfigPath()
	if path == "" {
		return nil, errors.New("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, errors.New("Failed to read config: " + err.Error())
	}

	return &cfg, nil
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
