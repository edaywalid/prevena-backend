package router

import (
	"github.com/edaywalid/pinktober-hackathon-backend/internal/di"
	"github.com/gorilla/mux"
)

type ProductRouter struct {
	container *di.Container
}

func NewProductRouter(container *di.Container) *ProductRouter {
	return &ProductRouter{container: container}
}

func (pr *ProductRouter) SetupRouter(router *mux.Router) {
	router.HandleFunc("/products", pr.container.Handlers.ProductHandler.GetProducts).Methods("GET")
	router.HandleFunc("/product/{barcode}", pr.container.Handlers.ProductHandler.GetProductByBarcode).Methods("GET")
	router.HandleFunc("/analyze", pr.container.Handlers.ProductHandler.AnalyzeIngredients).Methods("POST")
}
