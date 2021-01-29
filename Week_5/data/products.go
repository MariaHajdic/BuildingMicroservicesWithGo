package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
