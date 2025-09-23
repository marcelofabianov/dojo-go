package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	General  GeneralConfig  `mapstructure:"general"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Server   ServerConfig   `mapstructure:"server"`
	DB       DBConfig       `mapstructure:"db"`
	Password PasswordConfig `mapstructure:"password"`
}

type GeneralConfig struct {
	Env string `mapstructure:"env"`
	TZ  string `mapstructure:"tz"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type ServerConfig struct {
	API  APIConfig  `mapstructure:"api"`
	CORS CORSConfig `mapstructure:"cors"`
}

type APIConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	RateLimit    int           `mapstructure:"rate_limit"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	MaxBodySize  int           `mapstructure:"maxbodysize"`
}

type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowedorigins"`
	AllowedMethods   []string `mapstructure:"allowedmethods"`
	AllowedHeaders   []string `mapstructure:"allowedheaders"`
	ExposedHeaders   []string `mapstructure:"exposedheaders"`
	AllowCredentials bool     `mapstructure:"allowcredentials"`
}

type DBConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"name"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"maxopenconns"`
	MaxIdleConns    int           `mapstructure:"maxidleconns"`
	ConnMaxLifetime time.Duration `mapstructure:"connmaxlifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"connmaxidletime"`
	QueryTimeout    time.Duration `mapstructure:"querytimeout"`
	ExecTimeout     time.Duration `mapstructure:"exectimeout"`
}

type PasswordConfig struct {
	MinLength     int  `mapstructure:"min_length"`
	RequireNumber bool `mapstructure:"require_number"`
	RequireUpper  bool `mapstructure:"require_upper"`
	RequireLower  bool `mapstructure:"require_lower"`
	RequireSymbol bool `mapstructure:"require_symbol"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	v := viper.New()

	v.SetDefault("general.env", "development")
	v.SetDefault("general.tz", "UTC")
	v.SetDefault("logger.level", "info")
	v.SetDefault("server.api.host", "0.0.0.0")
	v.SetDefault("server.api.port", 8080)
	v.SetDefault("server.api.rate_limit", 100)
	v.SetDefault("server.api.read_timeout", "5s")
	v.SetDefault("server.api.write_timeout", "10s")
	v.SetDefault("server.api.idle_timeout", "120s")
	v.SetDefault("server.api.maxbodysize", 1048576)
	v.SetDefault("server.cors.allowedorigins", []string{"*"})
	v.SetDefault("server.cors.allowedmethods", []string{"GET", "POST"})
	v.SetDefault("server.cors.allowedheaders", []string{"Content-Type", "Authorization"})
	v.SetDefault("server.cors.exposedheaders", []string{})
	v.SetDefault("server.cors.allowcredentials", true)
	v.SetDefault("db.driver", "postgres")
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", 5432)
	v.SetDefault("db.user", "user")
	v.SetDefault("db.password", "password")
	v.SetDefault("db.name", "app")
	v.SetDefault("db.ssl_mode", "disable")
	v.SetDefault("db.maxopenconns", 10)
	v.SetDefault("db.maxidleconns", 5)
	v.SetDefault("db.connmaxlifetime", "1h")
	v.SetDefault("db.connmaxidletime", "10m")
	v.SetDefault("db.querytimeout", "5s")
	v.SetDefault("db.exectimeout", "3s")
	v.SetDefault("password.min_length", 12)
	v.SetDefault("password.require_number", true)
	v.SetDefault("password.require_upper", true)
	v.SetDefault("password.require_lower", true)
	v.SetDefault("password.require_symbol", false)

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
