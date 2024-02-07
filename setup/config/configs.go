package config

import (
	viper "github.com/spf13/viper"
)

type AppConfiguration interface {
	GetServerConfigs() ServerConfigs
	GetDatabaseConfigs() DatabaseConfigs
	GetMigrationConfigs() MigrationConfigs
}

type Configuration struct {
	ServerConfigs    ServerConfigs    `mapstructure:"server"`
	DatabaseConfigs  DatabaseConfigs  `mapstructure:"database"`
	MigrationConfigs MigrationConfigs `mapstructure:"migrationPaths"`
}

func (c *Configuration) GetServerConfigs() ServerConfigs {
	return c.ServerConfigs
}

func (c *Configuration) GetDatabaseConfigs() DatabaseConfigs {
	return c.DatabaseConfigs
}

func (c *Configuration) GetMigrationConfigs() MigrationConfigs {
	return c.MigrationConfigs
}

func InitConfigs(file string) Configuration {
	var configuration Configuration
	v := viper.New()
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		println("Error in reading configs: ", err.Error())
		return configuration
	}
	err = v.Unmarshal(&configuration)
	if err != nil {
		println("Error while unmarshalling configs: ", err.Error())
		return configuration
	}
	return configuration
}
