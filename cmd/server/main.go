package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/config"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/errors"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	cfg, err := config.Init()
	if err != nil {
		log.Fatalln(errors.ConfigInitError, err)
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(jsonHandler)

	server := server.NewServer(cfg.ServerHost, cfg.ServerPort, cfg.HashCashBits, logger)

	err = server.Start(ctx)
	if err != nil {
		log.Fatalln(errors.ServerStartError, err)
	}
}
