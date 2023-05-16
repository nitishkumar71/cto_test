package main

import (
	"fmt"
	ctoai "github.com/cto-ai/sdk-go/v2"
	"net/http"
)

func main() {
	client := ctoai.NewClient()
	client.Ux.Print("starting the server")

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
