package handlers

import (
	"go-microservices-tutorial/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to convert toJSON", http.StatusInternalServerError)
	}

}

//abstract out the conversion to JSON here
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to convert toJSON", http.StatusInternalServerError)
	}

}
