package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("Service: %s \nRequest Time: %s", os.Getenv("SERVICE"), time.Now().Format("Mon Jan _2 15:04:05 2006")))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
