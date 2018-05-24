package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/shutdown", shutdownHandler)

	log.Println("listening on port 80")
	log.Println(http.ListenAndServe(":80", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloWorldHandler requested")
	fmt.Fprintf(w, "hello from %s\n", os.Getenv("HOSTNAME"))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shutting down pod ", os.Getenv("HOSTNAME"))
	os.Exit(-1)
}
