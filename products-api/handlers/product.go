package handlers

import (
	"github.com/furkandemireleng/go-eventdriven/products-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		p.UpdateProduct(w, r)
		return
	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)

}
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	var err = lp.ToJson(w)

	if err != nil {
		http.Error(w, "Unable error", http.StatusInternalServerError)
	}
}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable unmarshall error", http.StatusInternalServerError)
	}

	data.AddProduct(prod)
}
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable unmarshall error", http.StatusInternalServerError)
	}

	data.UpdateProduct(prod)

}
