package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed docs/index.html docs/app.js docs/style.css docs/workbench.html docs/workbench.mjs docs/workbench.css docs/graph-view.mjs
var assets embed.FS

func respond(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
func app() http.Handler {
	mux := http.NewServeMux()
	files, _ := fs.Sub(assets, "docs")
	static := http.FileServer(http.FS(files))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "HEAD" {
			w.WriteHeader(405)
			return
		}
		if r.URL.Path != "/" && r.URL.Path != "/index.html" && r.URL.Path != "/app.js" && r.URL.Path != "/style.css" && r.URL.Path != "/workbench.html" && r.URL.Path != "/workbench.mjs" && r.URL.Path != "/workbench.css" && r.URL.Path != "/graph-view.mjs" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; connect-src 'self' https: http://localhost:* http://127.0.0.1:*; frame-ancestors 'none'; base-uri 'none'")
		static.ServeHTTP(w, r)
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		respond(w, 200, map[string]any{"status": "ok", "synthetic": true})
	})
	mux.HandleFunc("/api/demo-trades", func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) })

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		if o := r.Header.Get("Origin"); o != "" && !strings.EqualFold(o, "http://"+r.Host) {
			respond(w, 403, map[string]string{"error": "origin_denied"})
			return
		}
		mux.ServeHTTP(w, r)
	})
}
func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8090"
	}
	server := &http.Server{Addr: addr, Handler: app(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 30 * time.Second}
	log.Printf("Trade Agent Gate demo: http://%s", addr)
	log.Fatal(server.ListenAndServe())
}
