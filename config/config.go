package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	// Config - struct describes Application configs
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax     int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		DatabaseUri string `env-required:"true"                 env:"DATABASE_URI"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	err = ProcessArgs(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// ProcessArgs - processes and handles CLI arguments
func ProcessArgs(cfg *Config) error {
	// create flag set using `flag` package
	fset := flag.NewFlagSet("Example", flag.ContinueOnError)
	fset.StringVar(&cfg.PG.DatabaseUri, "d", cfg.PG.DatabaseUri, "PG connection url")

	// get config usage with wrapped flag usage
	fset.Usage = cleanenv.FUsage(fset.Output(), &cfg, nil, fset.Usage)

	err := fset.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	return nil
}
