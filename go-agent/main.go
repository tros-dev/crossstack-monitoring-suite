package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type Metrics struct {
	OS     string `json:"os"`
	CPUs   int    `json:"cpus"`
	Memory int64  `json:"memory"`
}

var logger = logrus.New()

func collectMetrics() Metrics {
	// Simulate concurrent data collection with goroutines
	var wg sync.WaitGroup
	var metrics Metrics

	wg.Add(3)

	go func() {
		defer wg.Done()
		metrics.OS = "linux" // Replace with actual system call if desired
		logger.WithField("metric", "os").Info("Collected OS info")
	}()

	go func() {
		defer wg.Done()
		metrics.CPUs = 16
		logger.WithField("metric", "cpus").Info("Collected CPU info")
	}()

	go func() {
		defer wg.Done()
		metrics.Memory = 17179869184
		logger.WithField("metric", "memory").Info("Collected memory info")
	}()

	wg.Wait()
	return metrics
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := collectMetrics()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: http.HandlerFunc(metricsHandler),
	}

	go func() {
		logger.Info("Starting Go agent on :9000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Failed to shutdown server: %v", err)
	}

	logger.Info("Server gracefully stopped")
}
