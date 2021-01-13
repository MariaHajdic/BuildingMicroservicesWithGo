package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running hello handler")

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", b)
	})

	log.Println("Starting the server")
	/* Listening for connections on all network interfaces on port 9090 */
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
