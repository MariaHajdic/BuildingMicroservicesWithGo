package data

import (
	"encoding/json"
	"io"
	"time"
)

/* Defines the structure for an API product */
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

/* Products is a collection of Product */
type Products []*Product

/* Method serialising a collection to JSON.
 * NewEncoder performs better than Unmarshal as it does not
 * buffers the output into an in-memory slice of bytes. It
 * reduces allocations and overheads of the service.
**/
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

/* A hard-coded example list of products.
**/
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Bubble tea",
		Description: "Some words",
		Price:       4.20,
		SKU:         "abedss",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "Delicious water",
		Description: "Some other words",
		Price:       0.40,
		SKU:         "aaa012",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
