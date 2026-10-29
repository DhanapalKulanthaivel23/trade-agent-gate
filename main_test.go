package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPublicBoundary(t *testing.T) {
	h := app()
	for _, path := range []string{"/", "/app.js", "/style.css", "/health", "/api/demo-trades"} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
		if w.Code != 200 {
			t.Fatal(path, w.Code)
		}
	}
	for _, path := range []string{"/AGENTS.md", "/go.mod", "/internal/trade/store.go", "/.env", "/api/transactions", "/mcp"} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
		if w.Code != 404 {
			t.Fatal("private path available", path, w.Code)
		}
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("POST", "/api/demo-trades", strings.NewReader("{}")))
	if w.Code != 405 {
		t.Fatal("unexpected mutation")
	}
}
