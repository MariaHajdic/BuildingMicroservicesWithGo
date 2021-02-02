package data

import "testing"

func TestCecksValidation(t *testing.T) {
	p := &Products{
		Name:  "name",
		Price: 1.00,
		SKU:   "abc-abd-dbe",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
