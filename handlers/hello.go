package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//interface
type Hello struct {
	l *log.Logger
}

// this function will return the Hello struct which also had the methdo below
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//creating a ServeHTTP method on the Hello struct
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//l is of type lo.Logger
	h.l.Println("Helloff World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello from Go, %s", d)
}
