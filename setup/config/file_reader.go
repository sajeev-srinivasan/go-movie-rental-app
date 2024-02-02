package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Port          int    `yaml:"port"`
		Host          string `yaml:"host"`
		Name          string `yaml:"name"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		Migrationpath string `yaml:"migrationpath"`
	} `yaml:"database"`
}

func GetConfig(cfg *Config, configFilePath string) {
	f, err := os.Open(configFilePath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
