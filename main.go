package main

import (
	"fmt"
	"net/http"
	"os"
)

// root
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// not root
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not fount\n"))
		return
	}

	// is root
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

// main function is run server
func main() {
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
