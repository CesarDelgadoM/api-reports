package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config general
type Config struct {
	Server   ServerConfig
	Services ServicesConf
	Cors     CorsConfig
	Postgres PostgresConfig
	Mongo    MongoConfig
}

type ServerConfig struct {
	Port string
}

// Services config
type ServicesConf struct {
	Producer ProducerServ
}

// Producer service
type ProducerServ struct {
	Url         string
	ContentType string
}

// Cors config
type CorsConfig struct {
	AllowCredentials bool
	AllowOrigins     string
	AllowHeaders     string
}

// Postgres config
type PostgresConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// Mongo config
type MongoConfig struct {
	URI      string
	User     string
	Password string
	DBName   string
}

func LoadConfig(fileName string) *viper.Viper {
	v := viper.New()

	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath("../config")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	return v
}

func ParseConfig(v *viper.Viper) *Config {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Panic(err)
	}

	return &c
}
