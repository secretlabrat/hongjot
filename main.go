package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/secretlabrat/hongjot/api/health"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/api/v1/health", health.Check())

	go func() { // comment here to simulate slow endpoint then Ctrl+C to stop the server
		if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil && err != http.ErrServerClosed {
			logger.Fatal("shutting down the server:", zap.Error(err))
		}
	}()

	logger.Info("Server is running on :%s", zap.String("port", os.Getenv("SERVER_PORT")))

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	sig, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-sig.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("shutting down the server:", zap.Error(err))
	}
	logger.Info("server shutdown gracefully")
}
