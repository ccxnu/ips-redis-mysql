package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Redis struct {
		Host string
		Port int
		Url  string
	}
	Database struct {
		User     string
		Password string
		Host     string
		Port     int
		Name     string
	}
	PollInterval int
	IpApiUrl     string
	TotalKeys    int64
	MatchPattern string
}

var AppConfig Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	viper.AutomaticEnv()

	AppConfig.Redis.Host = viper.GetString("REDIS_HOST")
	AppConfig.Redis.Port = viper.GetInt("REDIS_PORT")
	AppConfig.Redis.Url = viper.GetString("REDIS_URL")

	AppConfig.Database.User = viper.GetString("DB_USER")
	AppConfig.Database.Password = viper.GetString("DB_PASS")
	AppConfig.Database.Host = viper.GetString("DB_HOST")
	AppConfig.Database.Port = viper.GetInt("DB_PORT")
	AppConfig.Database.Name = viper.GetString("DB_NAME")

	AppConfig.IpApiUrl = viper.GetString("IP_API_URL")
	AppConfig.PollInterval = viper.GetInt("POLL_INTERVAL")
	AppConfig.TotalKeys = viper.GetInt64("TOTAL_KEYS")
	AppConfig.MatchPattern = viper.GetString("MATCH_PATTERN")
}
