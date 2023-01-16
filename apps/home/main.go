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

	van1 := vanity.Handler("m3o-dev.yaml")
	van2 := vanity.Handler("go-m3o-com.yaml")

	// new handler
	hdr := handler.New(srv, van1, van2)

	// register handler
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, hdr))

	// run on port 8080
	http.ListenAndServe(":8080", nil)
}
