package handlers

import (
	"net/http"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerHandler struct {
	config *config.Config
}

func NewSwaggerHandler(config *config.Config) *SwaggerHandler {
	return &SwaggerHandler{
		config: config,
	}
}
func (sh *SwaggerHandler) ServeYamlDocs(w http.ResponseWriter, r *http.Request) {
	yamlFile := "docs/swagger.yaml"
	w.Header().Set("Content-Type", "application/x-yaml")
	http.ServeFile(w, r, yamlFile)
}

func (sh *SwaggerHandler) ServeSwaggerUI() http.Handler {
	url := "http://localhost:" + sh.config.PORT + "/swagger/doc.yaml"
	if sh.config.IsProduction() {
		url = sh.config.PROD_URL + "/swagger/doc.yaml"
	}
	return httpSwagger.Handler(
		httpSwagger.URL(url),
	)
}
