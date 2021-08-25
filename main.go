package main

import (
	/* "fmt"
	"io/ioutil" */
	"go-microservices-tutorial/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//returns as a Hello struct?
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/goodbye", gb)
	sm.Handle("/", hh)

	http.ListenAndServe(":8080", sm)
}
