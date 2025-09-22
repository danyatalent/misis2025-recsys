package main

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogLevel        string `env:"LOG_LEVEL"          envDefault:"debug"`
	LogMaxSliceSize uint   `env:"LOG_MAX_SLICE_SIZE" envDefault:"4"`

	TwinWordAPIKey string `env:"TWINWORD_API_KEY" envDefault:""`
	TwinWordURL    string `env:"TWINWORD_URL"     envDefault:""`

	ApiLayerAPIKey string `env:"APILAYER_API_KEY" envDefault:""`
	ApiLayerURL    string `env:"APILAYER_URL"     envDefault:""`

	HTTPTimeout time.Duration `env:"HTTP_TIMEOUT" envDefault:"30s"`
}

func readConfig() (Config, error) {
	var config Config

	err := env.Parse(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return config, nil
}

func parseLogLevel(s string) (slog.Level, error) {
	var level slog.Level

	err := level.UnmarshalText([]byte(strings.ToUpper(s)))
	if err != nil {
		return slog.LevelInfo, fmt.Errorf("invalid log level: %w", err)
	}

	return level, nil
}
