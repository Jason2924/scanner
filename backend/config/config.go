package config

import "github.com/spf13/viper"

type Config struct {
	Mode        string            `mapstructure:"SERVER_MODE"`
	Port        string            `mapstructure:"SERVER_PORT"`
	OpenWeather ConfigOpenWeather `mapstructure:",squash"`
	Mysql       ConfigMysql       `mapstructure:",squash"`
	Redis       ConfigRedis       `mapstructure:",squash"`
}

type ConfigOpenWeather struct {
	ApiKey string `mapstructure:"OPEN_WEATHER_API_KEY"`
}

type ConfigMysql struct {
	Host         string `mapstructure:"MYSQL_HOST"`
	RootPassword string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	Name         string `mapstructure:"MYSQL_DATABASE"`
	Username     string `mapstructure:"MYSQL_USERNAME"`
	Password     string `mapstructure:"MYSQL_PASSWORD"`
	MigrateTable bool   `mapstructure:"MYSQL_MIGRATE_TABLE"`
	ImportData   bool   `mapstructure:"MYSQL_IMPORT_DATA"`
}

type ConfigRedis struct {
	Address  string `mapstructure:"REDIS_ADDRESS"`
	Username string `mapstructure:"REDIS_USERNAME"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

func Load(path, name, xten string) (*Config, error) {
	resp := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(xten)
	if xten == "env" {
		viper.AutomaticEnv()
	}
	if erro := viper.ReadInConfig(); erro != nil {
		return nil, erro
	}
	if erro := viper.Unmarshal(&resp); erro != nil {
		return nil, erro
	}
	return &resp, nil
}
