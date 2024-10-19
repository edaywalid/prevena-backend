package di

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/handlers"
)

type Container struct {
	Services     *Services
	Repositories *Repositories
	Handlers     *Handlers
	Config       *config.Config
}

type (
	Services     struct{}
	Repositories struct{}
	Handlers     struct {
		PingHandler *handlers.PingHandler
	}
)

func NewContainer() (*Container, error) {
	var container Container
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	container.Config = config
	container.Services = &Services{}
	container.Repositories = &Repositories{}
	container.Handlers = &Handlers{
		PingHandler: &handlers.PingHandler{},
	}

	return &container, nil
}
