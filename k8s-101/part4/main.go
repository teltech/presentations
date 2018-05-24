package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teltech/backbone"
	"github.com/teltech/logger"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorldHandler)
	mux.HandleFunc("/healthz", healthCheckHandler)
	log.Println("listening on port 80")

	backbone.ListenAndServe(":80", mux, logger.New())
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "response from V3\n")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
