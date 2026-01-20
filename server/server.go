package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//go:embed static/*
var staticFS embed.FS

var defaultTheme string

// Start starts the HTTP server
func Start(listen string, port int, theme string) {
	defaultTheme = theme
	if defaultTheme != "cyber" && defaultTheme != "terminal" {
		defaultTheme = "cyber"
	}

	mux := http.NewServeMux()

	// Static files
	mux.HandleFunc("GET /", handleIndex)
	mux.Handle("GET /static/", http.FileServerFS(staticFS))

	// API v1 endpoints
	mux.HandleFunc("GET /api/v1/generate", handleGenerate)
	mux.HandleFunc("GET /api/v1/batch", handleBatch)
	mux.HandleFunc("GET /api/v1/export", handleExport)

	addr := fmt.Sprintf("%s:%d", listen, port)
	log.Infof("Server starting at http://%s", addr)
	log.Infof("Default theme: %s", defaultTheme)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFS(staticFS, "static/index.html")
	if err != nil {
		log.Errorf("Failed to parse template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		DefaultTheme string
	}{
		DefaultTheme: defaultTheme,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		log.Errorf("Failed to execute template: %v", err)
	}
}

func parseCount(r *http.Request) int {
	countStr := r.URL.Query().Get("count")
	if countStr == "" {
		return 1
	}
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		return 1
	}
	if count > 1000 {
		return 1000
	}
	return count
}
