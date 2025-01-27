package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig AppConfig
	DbConfig  DbConfig
}

type AppConfig struct {
	AppName  string
	AppPort  string
	BasePath string
}

type DbConfig struct {
	DBName          string
	DBConnectionUri string
	Timeout         int
	DBNamePrefix    string
}

var ServerConf *Config

func LoadConfiguration() {
	viper.SetDefault("APP_NAME", "go-and-mongo")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("BASE_PATH", "/go-and-mongo")

	viper.SetDefault("DB_NAME", "go-and-mongo")
	viper.SetDefault("DB_URI", "mongodb+srv://edward505:Skateer.55@cluster0.fgxd6os.mongodb.net/")
	viper.SetDefault("DB_TIMEOUT", 10)
	viper.SetDefault("DB_NAME_PREFIX", "tst")

	readConfigFile("../your-path-file")

	ServerConf = &Config{
		AppConfig: AppConfig{
			AppName:  viper.GetString("APP_NAME"),
			AppPort:  viper.GetString("APP_PORT"),
			BasePath: viper.GetString("BASE_PATH"),
		},
		DbConfig: DbConfig{
			DBName:          viper.GetString("DB_NAME"),
			DBConnectionUri: viper.GetString("DB_URI"),
			Timeout:         viper.GetInt("DB_TIMEOUT"),
			DBNamePrefix:    viper.GetString("DB_NAME_PREFIX"),
		},
	}
}

func readConfigFile(path string) {
	viper.AutomaticEnv()
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: No config file found, using default values")
	}
}
