package router

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/gorilla/mux"
)

type SeedRouter struct {
	container *di.Container
}

func NewSeedRouter(container *di.Container) *SeedRouter {
	return &SeedRouter{container: container}
}

func (pr *SeedRouter) SetupRouter(router *mux.Router) {
	router.HandleFunc("/seed", pr.container.Handlers.SeedHandler.Seed).Methods("GET")
}
