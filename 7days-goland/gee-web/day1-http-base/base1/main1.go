package main

import (
	"fmt"
	"log"
	"net/http"
)

//$ curl http://localhost:9999/
//URL.Path = "/"
//$ curl http://localhost:9999/hello
//Header["Accept"] = ["*/*"]
//Header["User-Agent"] = ["curl/7.54.0"]

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// handler echoes r.URL.Path
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// handler echoes r.URL.Header
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(":9999", nil))
}
