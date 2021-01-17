package handlers

import (
	"Week_3/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/* The main entry point for the handler, satisfies the
 * http.Handler interface.
**/
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Catch all - if no method is satisfied, return an error.
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

/* Returns the products from the data store.
**/
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetching the products from the data store.
	lp := data.GetProducts()

	// Serializing the list to JSON.
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
