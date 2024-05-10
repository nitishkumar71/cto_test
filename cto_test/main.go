package main

import (
	"fmt"
	"log"
	"net/http"

	ctoai "github.com/cto-ai/sdk-go/v2"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func main() {
	client := ctoai.NewClient()
	client.Ux.Print("starting the server")

	event := map[string]interface{}{
		"event_name":   "deployment",
		"event_action": "success",
		"branch":       "fix_sdk_ops_run",
		"repo":         "nitishkumar71/cto_test",
		"environment":  "production",
	}

	err := client.Sdk.Track([]string{}, "", event)
	if err != nil {
		panic(err)
	}

	fmt.Println("DUmmy Line")
	http.HandleFunc("/", logging(HelloServer))
	http.HandleFunc("/shorten", logging(ShortURLHandler))
	http.HandleFunc("/redirect/", logging(redirectURLHandler))
	http.HandleFunc("/form", logging(formHandle))
	http.HandleFunc("/json-encode/", logging(jsonEncode))
	http.HandleFunc("/json-decode/", logging(jsonDecode))
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
