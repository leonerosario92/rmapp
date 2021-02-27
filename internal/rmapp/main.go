package main

import (
	"github.com/go-chi/chi"
	"github.com/leonerosario92/rmapp/internal/common/server"
	"github.com/leonerosario92/rmapp/internal/rmapp/ports"
	"net/http"
)

func main() {

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})
}
