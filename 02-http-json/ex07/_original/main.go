package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type middleware func(http.Handler) http.Handler

func chain(handler http.Handler, middlewares ...middleware) http.Handler {
	// TODO: implement
	return handler
}

func withServerHeader(next http.Handler) http.Handler {
	// TODO: implement
	return next
}

func withRecovery(next http.Handler) http.Handler {
	_ = json.NewEncoder
	// TODO: implement
	return next
}

func newHandler() http.Handler {
	mux := http.NewServeMux()

	// TODO: register routes using stdlib patterns:
	// - GET /health
	// - GET /ready
	// - POST /echo
	// - GET /panic

	return chain(mux, withServerHeader, withRecovery)
}

func newServer(addr string, handler http.Handler) *http.Server {
	_ = time.Second
	// TODO: implement explicit server configuration and timeouts
	return &http.Server{}
}

func serve(ctx context.Context, srv *http.Server) error {
	_ = errors.Is
	_ = context.WithTimeout
	// TODO: implement graceful shutdown
	// Requirements:
	// - call srv.Shutdown with a timeout when ctx is canceled
	// - treat http.ErrServerClosed as a normal shutdown
	return nil
}

func main() {
	handler := newHandler()
	server := newServer(":8080", handler)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	_ = serve(ctx, server)
}
