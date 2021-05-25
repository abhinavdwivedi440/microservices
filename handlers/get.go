package handlers

import (
	"net/http"

	"github.com/abhinavdwivedi440/microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET")
	w.Header().Set("Content-Type", "application/json")

	// fetch the products from the datastore
	prodList := data.GetProducts()

	// serialize the list to JSON
	err := prodList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}
