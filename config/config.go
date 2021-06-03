package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	once sync.Once

	Cors struct {
		AllowedHeaders []string `yaml:"allowed_headers"`
		AllowedOrigins []string `yaml:"allowed_origins"`
		AllowedMethods []string `yaml:"allowed_methods"`
	} `yaml:"cors"`

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	MongoDB struct {
		Uri string `yaml:"uri"`
		Db  string `yaml:"db"`
	} `yaml:"mongo_db"`

	ClientId string `yaml:"client_id"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func Init() *Config {
	config := Config{}
	config.once.Do(func() { config.init("config.yaml") })
	return &config
}

func (config *Config) init(file string) {
	f, err := os.Open(file)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		processError(err)
	}
}

func (config *Config) GetClientId() string {
	return config.ClientId
}

func (config *Config) GetDbName() string {
	return config.MongoDB.Db
}

func (config *Config) GetMongoDbUri() string {
	return config.MongoDB.Uri
}

func (config *Config) AllowedHeaders() []string {
	return config.Cors.AllowedHeaders
}

func (config *Config) AllowedOrigins() []string {
	return config.Cors.AllowedOrigins
}

func (config *Config) AllowedMethods() []string {
	return config.Cors.AllowedMethods
}

func (config *Config) GetHost() string {
	return config.Server.Host
}

func (config *Config) GetPort() string {
	return config.Server.Port
}
