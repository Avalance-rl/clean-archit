package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"net"
	"os"
	"sync"
)

var (
	configInstance *Config
	once           sync.Once
)

type Config struct {
	Env string `yaml:"env" env-default:"prod"`
	HTTPServer
	DB    Database `yaml:"db"`
	Redis Redis    `yaml:"redis"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
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
		c.DB.Host, c.DB.Port, c.DB.SSLMode, c.DB.User, c.DB.Password, c.DB.Name,
	)
}

func (c *Config) BuildServerAddress() string {
	return net.JoinHostPort(c.HTTPServer.Host, c.HTTPServer.Port)
}
