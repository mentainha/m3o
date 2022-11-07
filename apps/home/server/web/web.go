// Package web includes the web frontend
package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed html/*
var html embed.FS

// Handle registers the web handler
func Register(mux *http.ServeMux) {
	htmlContent, err := fs.Sub(html, "html")
	if err != nil {
		log.Fatal(err)
	}
	// serve the html directory
	mux.Handle("/", http.FileServer(http.FS(htmlContent)))
}
