package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
		cowFooter := `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hi there %s " + cowFooter + "\n\n" + host, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
