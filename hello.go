//hello bitch

package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my nightmare</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact</h1><a href=\"mailto:treighton@gmail.com\">Message me bitch</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404<br />UM WTF GET OUTTA HERE</h1>")
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":3000", nil)
}
