package router

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/gorilla/mux"
)

type PingRouter struct {
	container *di.Container
}

func NewPingRouter(container *di.Container) *PingRouter {
	return &PingRouter{container: container}
}

func (pr *PingRouter) SetupRouter(router *mux.Router) {
	router.HandleFunc("/ping", pr.container.Handlers.PingHandler.Ping).Methods("GET")
}
