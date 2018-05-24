package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received a request")
		resp, err := http.Get("http://hello-service")
		if err != nil {
			log.Println("Error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		fmt.Fprintf(w, "Results: %s\n", string(bodyBytes))
	})

	log.Println("listening on port 80")
	log.Println(http.ListenAndServe(":80", nil))
}
