package router

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/gorilla/mux"
)

func SetupRouter(container *di.Container) *mux.Router {
	r := mux.NewRouter()
	NewPingRouter(container).SetupRouter(r)
	NewSwaggerRouter(container).SetupRouter(r)
	return r
}
