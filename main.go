package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, os.Getenv("SERVICE"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
