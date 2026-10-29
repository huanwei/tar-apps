package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/huanwei/tar-apps/internal/farewell"
	"github.com/huanwei/tar-apps/internal/greeting"
	"github.com/huanwei/tar-apps/internal/version"
)

func main() {
	addr := ":8080"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealthz)
	mux.HandleFunc("GET /readyz", handleReadyz)
	mux.HandleFunc("GET /hello", handleHello)
	mux.HandleFunc("GET /farewell", handleFarewell)
	mux.HandleFunc("GET /version", handleVersion)

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok\n"))
}

func handleReadyz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	writeJSON(w, map[string]string{"message": greeting.Hello(name)})
}

func handleFarewell(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	writeJSON(w, map[string]string{"message": farewell.Farewell(name)})
}

func handleVersion(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{"version": version.Version})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "encode response", http.StatusInternalServerError)
	}
}
