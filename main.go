package main

import (
	/* "fmt"
	"io/ioutil" */
	"go-microservices-tutorial/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//each handler is returning the struct, which is like a route
	//http handlers have a ServeHTTP method
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)
	productsHandler := handlers.NewProducts(l)

	//server multiplexer is the handler
	sm := http.NewServeMux()
	sm.Handle("/goodbye", gb)
	sm.Handle("/", hh)
	sm.Handle("/products", productsHandler)

	//this is to prevent stuff like dos
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()

}
