package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

var (
	configInstance *Config
	once           sync.Once
)

type Config struct {
	Env   string   `yaml:"env" env-default:"prod"`
	DB    Database `yaml:"db"`
	Redis Redis    `yaml:"redis"`
}

type Database struct {
	host     string `yaml:"host"`
	port     string `yaml:"port"`
	sslMode  string `yaml:"sslmode"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
	name     string `yaml:"name"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		configInstance.mustLoad()
	})
	return configInstance
}

func (c *Config) mustLoad() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	if err := cleanenv.ReadConfig(path, configInstance); err != nil {
		panic("failed to read config: " + err.Error())
	}

}

func (c *Config) BuildPostgresDns() string {
	return fmt.Sprintf(
		"host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		c.DB.host, c.DB.port, c.DB.sslMode, c.DB.user, c.DB.password, c.DB.name,
	)
}
