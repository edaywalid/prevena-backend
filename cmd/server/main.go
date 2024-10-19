package main

import (
	"net/http"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/router"
	"github.com/edaywalid/pinktober-hackathon-backend/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	container, err := di.NewContainer()
	if err != nil {
		log.LogError().Msgf("Error creating container: %v", err)
		return
	}
	log.LogInfo().Msg("Container created successfully")

	log.LogInfo().Msg("Setting up router")
	router := router.SetupRouter(container)
	log.LogInfo().Msg("Router setup successfully")

	log.LogInfo().Msg("Starting server")
	log.LogInfo().Msgf("Server started on port %s", container.Config.PORT)

	if err := http.ListenAndServe(":"+container.Config.PORT, router); err != nil {
		log.LogError().Msgf("Error starting server: %v", err)
	}

}
