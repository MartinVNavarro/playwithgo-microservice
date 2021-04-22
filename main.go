package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/martinvnavarro/playwithgo-microservice/handler"
)

func main() {
	l := log.New(os.Stdout, "test-microservice", log.LstdFlags)

	hh := handler.NewIndex(l)
	mm := handler.NewMisc(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/misc", mm)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// For graceful shutdown
	sig := <-sigChan
	l.Println("Signal received, shutting down...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
