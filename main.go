// Command mock-api is a minimal HTTP service that mirrors the shape of the
// CCF api in miniature, for developing shared CI and release workflows.
package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/compliance-framework/mock-api/pkg/version"
)

// buildVersion is the version of this binary. The Dockerfile and Makefile set
// it with -ldflags "-X main.buildVersion=<VERSION>".
var buildVersion = "dev"

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	return mux
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func main() {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           newMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("mock-api %s (lib %s) listening on %s", buildVersion, version.Version, srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
