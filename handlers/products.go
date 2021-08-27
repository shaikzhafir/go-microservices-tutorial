package handlers

import (
	"encoding/json"
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

	//convert products into JSON
	data, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "unable to marshal JSON", http.StatusInternalServerError)
	}

	rw.Write(data)

}
