package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/silveiralexf/goflat/pkg/server"
)

func main() {
	slog.Info("starting server", "host", hostAddr)

	s := server.New(hostAddr, time.Second*30)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	ctx, stopCtx := context.WithCancel(context.Background())

	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				slog.ErrorContext(ctx, "graceful shutdown timed out.. forcing exit.")
				os.Exit(http.StatusGatewayTimeout)
			}
		}()

		err := s.Shutdown(shutdownCtx)
		if err != nil {
			slog.ErrorContext(context.Background(), err.Error())
			os.Exit(http.StatusRequestTimeout)
		}
		slog.Debug("graceful shutdown initiated", "host", hostAddr)
		cancel()
		stopCtx()
	}()

	err := s.Start()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}

	<-ctx.Done()
}

var (
	hostAddr string
	logLevel int64
)

func init() {
	flag.StringVar(
		&hostAddr,
		"hostAddr",
		"localhost:3000",
		"Informs host and port, e.g: localhost:3000",
	)
	flag.Int64Var(
		&logLevel,
		"logLevel",
		int64(slog.LevelError),
		"Sets loglevel: Debug=-4, Info=0, Warn=4, Error=8",
	)
	flag.Parse()

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				Level: slog.Level(logLevel),
			},
		),
	)
	slog.SetDefault(logger)
}
