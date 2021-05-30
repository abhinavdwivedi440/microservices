package handlers

import (
	"net/http"

	"github.com/abhinavdwivedi440/microservices/data"
)

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
