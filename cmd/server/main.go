package main

import (
	"canvas/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var release string

func main() {
	os.Exit(start())
}

func start() int {
	logEnv := getStringOrDefault("LOG_ENV", "development")
	log, err := createLogger(logEnv)
	if err != nil {
		fmt.Println("Error setting up logger:", err)
		return 1
	}

	log = log.With(zap.String("release", release))

	defer func() {
		// Flushes buffer, if any
		_ = log.Sync()
	}()

	host := getStringOrDefault("HOST", "localhost")
	port := getIntOrDefault("PORT", 8080)

	s := server.New(server.Options{
		Host: host,
		Port: port,
		Log:  log,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.Start(); err != nil {
			log.Info("error starting server", zap.Error(err))
			return err
		}
		return nil
	})

	<-ctx.Done()
	eg.Go(func() error {
		if err := s.Stop(); err != nil {
			log.Info("error stopping server", zap.Error(err))
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 1
	}

	return 0
}

func createLogger(env string) (*zap.Logger, error) {
	switch env {
	case "production":
		return zap.NewProduction()
	case "development":
		return zap.NewDevelopment()
	default:
		return zap.NewNop(), nil
	}
}

func getStringOrDefault(name, defaultValue string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return v
}

func getIntOrDefault(name string, defaultValue int) int {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	vAsInt, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return vAsInt
}
