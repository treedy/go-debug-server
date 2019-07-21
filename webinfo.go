package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from %s!\n", hostname)
	query := r.URL.Query()
	if query.Get("cookies") != "" {
		fmt.Fprintf(w, "Cookies:  %s\n", r.Cookies())
	}
	if query.Get("headers") != "" {
		fmt.Fprintf(w, "Headers:  %s\n", r.Header)
	}
	if query.Get("time") != "" {
		fmt.Fprintf(w, "Server time: %s", time.Now())
	}
}

func main() {
	http.HandleFunc("/", greet)
	fmt.Println("Listening on port 8080 for connections...")
	http.ListenAndServe(":8080", nil)
}
