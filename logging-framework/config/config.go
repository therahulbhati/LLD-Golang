package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const (
	ConsoleSinkType  = "console"
	FileSinkType     = "file"
	DatabaseSinkType = "database"
	InMemorySinkType = "inmemory"
)

type Config struct {
	TimeFormat string            `yaml:"time_format"`
	LogLevels  map[string]string `yaml:"log_levels"`
	FilePath   string            `yaml:"file_path"`
	DSN        string            `yaml:"dsn"`
}

func LoadConfig(file string) (*Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
