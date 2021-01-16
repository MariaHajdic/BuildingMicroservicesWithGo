package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

/* Implements go http.Handler interface */
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle goodbye request")
	/* Writing the response */
	fmt.Fprintf(rw, "Goodbye")
}
