//hello bitch

package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my nightmare</h1>")
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my shit</h1>")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/other", otherHandler)
	http.ListenAndServe(":3000", nil)
}
