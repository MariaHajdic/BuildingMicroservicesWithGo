package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

/* Implements go http.Handler interface */
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle hello requests")

	/* Reading the request body */
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}
	/* Writing the response */
	fmt.Fprintf(rw, "Hello %s", b)
}
