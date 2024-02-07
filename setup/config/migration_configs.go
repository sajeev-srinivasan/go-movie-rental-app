package config

type MigrationConfigs struct {
	DevPath  string `mapstructure:"dev"`
	TestPath string `mapstructure:"test"`
}
