package main

import (
	"Week_3/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address to the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "Week_3", log.LstdFlags)

	/* Creating new handlers */
	ph := handlers.NewProducts(l)

	/* Creating new serve mux and registering the handlers */
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	/* Creating a new server */
	s := http.Server{
		Addr:         *bindAddress,      // configuring the bind address
		Handler:      sm,                // setting the default handler
		ErrorLog:     l,                 // setting the logger for the server
		ReadTimeout:  5 * time.Second,   // timeout for reading client request
		WriteTimeout: 10 * time.Second,  // timeout for writing response
		IdleTimeout:  120 * time.Second, // timeout for connections using TCP Keep-Alive
	}

	/* Starting the server */
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	/* Trapping sigterm or interrupt */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	/* Blocking until a signal is received */
	sig := <-c
	log.Println("Got signal:", sig)

	/* Gracefully sutting down, waiting 30 Sec for operations to complete */
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
