package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	"github.com/edaywalid/pinktober-hackathon-backend/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerHandler struct {
	config *config.Config
	logger *logger.MyLogger
}

func NewSwaggerHandler(
	config *config.Config,
	logger *logger.MyLogger,
) *SwaggerHandler {
	return &SwaggerHandler{
		config: config,
		logger: logger,
	}
}
func (sh *SwaggerHandler) ServeYamlDocs(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	yamlFile := filepath.Join(cwd, "docs", "swagger.yaml")
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		sh.logger.LogError().Msg("Swagger file not found")
	}
	sh.logger.LogInfo().Msg("Serving swagger file")

	w.Header().Set("Content-Type", "application/x-yaml")
	http.ServeFile(w, r, yamlFile)
}

func (sh *SwaggerHandler) ServeSwaggerUI() http.Handler {
	url := "http://localhost:" + sh.config.PORT + "/swagger/doc.yaml"
	if sh.config.IsProduction() {
		url = sh.config.DOCS_URL
	}

	return httpSwagger.Handler(
		httpSwagger.URL(url),
	)
}
