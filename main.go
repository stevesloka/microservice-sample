package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	argListenPort = flag.Int("listen-port", 8080, "port to have API listen")
)

// Default (GET "/")
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Service: %s \nRequest Time: %s", os.Getenv("SERVICE"), time.Now().Format("Mon Jan _2 15:04:05 2006")))
}

// Default (GET "/patient_document/id")
func patientDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	docid := vars["docid"]
	fmt.Fprintf(w, fmt.Sprintf("!PATIENT DOCUMENT!\nDocId: %s\nService: %s \nRequest Time: %s", docid, os.Getenv("SERVICE"), time.Now().Format("Mon Jan _2 15:04:05 2006")))
}

func main() {
	flag.Parse()
	log.Println("--- [Sample Go Microservice is up and running!] ---", time.Now())

	// Configure router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/patient_document/{docid}", patientDocument)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *argListenPort), router))
}
