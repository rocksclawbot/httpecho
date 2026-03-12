package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type EchoResponse struct {
	Method     string              `json:"method"`
	Path       string              `json:"path"`
	Query      string              `json:"query,omitempty"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body,omitempty"`
	RemoteAddr string              `json:"remote_addr"`
	Timestamp  string              `json:"timestamp"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp := EchoResponse{
		Method:     r.Method,
		Path:       r.URL.Path,
		Query:      r.URL.RawQuery,
		Headers:    r.Header,
		Body:       string(body),
		RemoteAddr: r.RemoteAddr,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)

	fmt.Printf("🔊 httpecho listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
