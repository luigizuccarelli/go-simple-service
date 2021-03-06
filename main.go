package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/microlib/simple"
)

var (
	logger *simple.Logger
)

func startHttpServer(port string, logger *simple.Logger) *http.Server {
	srv := &http.Server{Addr: ":" + port}

	http.HandleFunc("/api/v1/service", func(w http.ResponseWriter, r *http.Request) {
		SimpleHandler(w, r, logger)
	})
	http.HandleFunc("/isalive", IsAlive)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Httpserver: ListenAndServe() error: " + err.Error())
		}
	}()

	return srv
}

func main() {

	logger = &simple.Logger{Level: "trace"}

	srv := startHttpServer(os.Args[1], logger)
	logger.Info("Starting server on port " + srv.Addr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-c
			switch s {
			case syscall.SIGHUP:
				exit_chan <- 0
			case syscall.SIGINT:
				exit_chan <- 0
			case syscall.SIGTERM:
				exit_chan <- 0
			case syscall.SIGQUIT:
				exit_chan <- 0
			default:
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan

	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	logger.Info("Server shutdown successfully")
	os.Exit(code)
}
