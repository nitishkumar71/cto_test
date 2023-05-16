package main

import (
	"fmt"
	"net/http"

	ctoai "github.com/cto-ai/sdk-go/v2"
)

func main() {
	client := ctoai.NewClient()
	client.Ux.Print("starting the server")

	event := map[string]interface{}{
		"event_name":   "deployment",
		"event_action": "succeeded",
		"branch":       "main",
		"repo":         "cto_test",
		"environment":  "production",
	}

	err := client.Sdk.Track([]string{}, "", event)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
