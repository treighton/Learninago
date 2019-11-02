//hello bitch

package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my nightmare</h1>")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":3000", nil)
}
