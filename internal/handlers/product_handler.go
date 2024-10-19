package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/dto"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/services"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}

	paginatedResponse, err := h.service.GetPaginatedProducts(r.Context(), page, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedResponse)
}

func (h *ProductHandler) GetProductByBarcode(w http.ResponseWriter, r *http.Request) {
	barcode := mux.Vars(r)["barcode"]
	product, err := h.service.GetProductByBarcode(r.Context(), barcode)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) AnalyzeIngredients(w http.ResponseWriter, r *http.Request) {
	var req dto.AnalyzeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	analyzedProduct, err := h.service.AnalyzeIngredients(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analyzedProduct)
}
