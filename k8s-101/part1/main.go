package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Println("listening on port 80")
	http.ListenAndServe(":80", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from %s\n", os.Getenv("HOSTNAME"))
}
