package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/client"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/config"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/errors"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Init()
	if err != nil {
		log.Fatalln(errors.ConfigInitError, err)
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(jsonHandler)
	client := client.NewClient(cfg.ClientHost, cfg.ClientPort, cfg.RepeatedCount, cfg.HashCashBits, logger)

	err = client.Run(ctx)
	if err != nil {
		log.Fatalln(errors.ClientStartError, err)
	}
}
