package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "tea",
		Price: 1,
		SKU:   "abc-def-efg",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
