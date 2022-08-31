package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"m3o.dev/cmd/api-gateway/handler"
)

func main() {
	// logging handler
	handler := handlers.CombinedLoggingHandler(os.Stdout, handler.New())
	// register the proxy handler
	http.Handle("/", handler)
	// run on port 8080
	http.ListenAndServe(":8080", nil)
}
