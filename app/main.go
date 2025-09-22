package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/danyatalent/misis2025-recsys/pkg/adapters/apilayer"
	"github.com/danyatalent/misis2025-recsys/pkg/adapters/twinword"
	"github.com/danyatalent/misis2025-recsys/pkg/presentation"
	"github.com/danyatalent/misis2025-recsys/pkg/usecase/analysis"
	"github.com/golang-cz/devslog"
)

func main() {
	ctx := context.Background()

	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	err = initLogger(cfg)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	err = run(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run(ctx context.Context, cfg Config) error {
	apiLayerClient, err := apilayer.New(apilayer.Config{
		BaseURL: cfg.ApiLayerURL,
		Key:     cfg.ApiLayerAPIKey,
		Timeout: cfg.HTTPTimeout,
	})
	if err != nil {
		return err
	}

	twinwordClient, err := twinword.New(twinword.Config{
		BaseURL: cfg.TwinWordURL,
		Key:     cfg.TwinWordAPIKey,
	})
	if err != nil {
		return err
	}

	uc, err := analysis.New(apiLayerClient, twinwordClient)
	if err != nil {
		return err
	}

	gui, err := presentation.New(uc)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)

	gui.Run(ctx, cancel)

	return nil
}

func initLogger(cfg Config) error {
	logLvl, err := parseLogLevel(cfg.LogLevel)
	if err != nil {
		return err
	}

	slogOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     logLvl,
	}

	opts := &devslog.Options{
		HandlerOptions:    slogOpts,
		MaxSlicePrintSize: cfg.LogMaxSliceSize,
		SortKeys:          true,
		NewLineAfterLog:   true,
		StringerFormatter: true,
	}

	logger := slog.New(devslog.NewHandler(os.Stdout, opts))

	slog.SetDefault(logger)

	return nil
}
