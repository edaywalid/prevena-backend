package router

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/gorilla/mux"
)

type SwaggerRouter struct {
	container *di.Container
}

func NewSwaggerRouter(container *di.Container) *SwaggerRouter {
	return &SwaggerRouter{container: container}
}

func (sr *SwaggerRouter) SetupRouter(router *mux.Router) {
	router.HandleFunc("/swagger/doc.yaml", sr.container.Handlers.SwaggerHandler.ServeYamlDocs).Methods("GET")
	router.PathPrefix("/swagger/").Handler(sr.container.Handlers.SwaggerHandler.ServeSwaggerUI())
}
