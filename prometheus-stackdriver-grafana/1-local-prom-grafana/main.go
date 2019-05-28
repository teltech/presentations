package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	listenAddr := os.Getenv("ADDR")
	if listenAddr == "" {
		listenAddr = ":5000"
	}

	workerName := os.Getenv("WORKER_NAME")
	if workerName == "" {
		workerName = "worker_a"
	}

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	unitOfWorkCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "dummy_app",
			Name:      "unit_of_work_total",
			Help:      "Total number of tasks completed.",
		},
		[]string{"worker_name"},
	)

	if err := prometheus.Register(unitOfWorkCounter); err != nil {
		logger.Println("unitOfWorkCounter not registered:", err)
	} else {
		logger.Println("unitOfWorkCounter registered.")
	}

	router := http.NewServeMux()
	router.Handle("/", prometheus.InstrumentHandler("index", (index(unitOfWorkCounter, workerName))))
	router.Handle("/metrics", prometheus.Handler())

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      (router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func index(workCounter *prometheus.CounterVec, workerName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// adding fake latency
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		workCounter.WithLabelValues(workerName).Inc()

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, World!")
	})
}
