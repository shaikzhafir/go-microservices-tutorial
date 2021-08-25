package handlers

import (
	//"fmt"
	//"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	//"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye world, nice knowing u")
	json.NewEncoder(rw).Encode(map[string]string{"status": "OK"})
}
