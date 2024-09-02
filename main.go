package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SUKSESS BROOO")
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT") // Render membutuhkan penggunaan PORT dari env var
	if port == "" {
		port = "8080" // Port default jika tidak ada env var
	}
	http.ListenAndServe(":"+port, nil)
}
