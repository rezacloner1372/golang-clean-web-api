package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
}

type ServerConfig struct {
	Port         string
	ExternalPort string
	runMode      string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yaml")
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())

	}
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error parse config file: %s", err.Error())

	}
	return cfg
}

// convert to config struct
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

// read  config file and convert to viper struct
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to loade config file: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// get config file path based on env
func getConfigPath(env string) string {
	if env == "docker" {
		return "../config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "../config/config-development"
	}
}
