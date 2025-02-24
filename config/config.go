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
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Log    `yaml:"logger"`
		PG     `yaml:"postgres"`
		Worker `yaml:"worker"`
	}

	App struct {
		Name      string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version   string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Env       string `env-required:"true"                env:"ENV"`
		JWTSecret string `env-required:"true"                env:"JWT_SECRET"`
	}

	HTTP struct {
		Port        string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Addr        string `env-required:"true"             env:"RUN_ADDRESS"`
		AccrualAddr string `env-required:"true"             env:"ACCRUAL_SYSTEM_ADDRESS"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax     int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		DatabaseUri string `env-required:"true"                 env:"DATABASE_URI"`
	}

	Worker struct {
		Poling int `env-required:"true" yaml:"poling"`
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
	fset.StringVar(&cfg.HTTP.Addr, "a", cfg.HTTP.Addr, "Http host and port")
	fset.StringVar(&cfg.HTTP.AccrualAddr, "r", cfg.HTTP.AccrualAddr, "Accrual system address")

	// get config usage with wrapped flag usage
	fset.Usage = cleanenv.FUsage(fset.Output(), &cfg, nil, fset.Usage)

	err := fset.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	return nil
}
