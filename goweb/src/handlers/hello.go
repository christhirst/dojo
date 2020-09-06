package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		h.l.Println("Hello Word")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error", http.StatusBadRequest)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})
}
