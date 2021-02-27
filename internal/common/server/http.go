package server

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()

	rootRouter := chi.NewRouter()

	rootRouter.Mount("/api", createHandler(apiRouter))

	http.ListenAndServe(addr, rootRouter)
}
