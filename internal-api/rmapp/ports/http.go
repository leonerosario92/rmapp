package ports

import (
	"net/http"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) GetGroups(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h HttpServer) CreateGroup(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
