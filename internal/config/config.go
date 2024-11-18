package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string     `yaml:"env"`
	Database Database   `yaml:"database"`
	Server   GrpcServer `yaml:"server"`
}

type GrpcServer struct {
	Address string        `yaml:"address"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
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

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
