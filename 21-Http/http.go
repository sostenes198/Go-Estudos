package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP em GO")

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
