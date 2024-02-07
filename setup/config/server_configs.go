package config

type ServerConfigs struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
