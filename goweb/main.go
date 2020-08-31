package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello Word")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error", http.StatusBadRequest)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.ListenAndServe(":9090", nil)
}
