package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"m3o.dev/apps/home/handler"
	"m3o.dev/apps/home/server"
	"m3o.dev/apps/home/vanity"
)

func main() {
	// the home server
	srv := server.New()

	van := vanity.Handler()

	// new handler
	hdr := handler.New(srv, van)

	// register handler
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, hdr))

	// run on port 8080
	http.ListenAndServe(":8080", nil)
}
