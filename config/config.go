package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		AUTH `yaml:"authentication"`
		RMQ  `yaml:"rabbitmq"`
		MDB  `yaml:"mongodb"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port           string   `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		AllowedOrigins []string `env-required:"true" yaml:"allowed-origins" env:"ALLOWED_ORIGINS"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	//keycloak auth
	AUTH struct {
		Secret string `env-required:"true" yaml:"secret" env:"AUTHENTICATION_SECRET"`
	}

	// RMQ -.
	RMQ struct {
		URL   string `env-required:"true" yaml:"url"    env:"RABBITMQ_URL"`
		QUEUE string `env-required:"true" yaml:"queue"  env:"RABBITMQ_QUEUE"`
	}

	//mongodb
	MDB struct {
		Username string `env-required:"true" yaml:"username" env:"MDB_USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"MDB_PASSWORD"`
		Cluster  string `env-required:"true" yaml:"cluster"  env:"MDB_CLUSTER"`
		Database string `env-required:"true" yaml:"database" env:"MDB_DATABASE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// NewConfig returns app config.
func NewIntergrationTestConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
