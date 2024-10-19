package services

import (
	"context"
	"math"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/dto"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/repositories"
)

type ProductService struct {
	product_repo    *repositories.ProductRepository
	ingredient_repo *repositories.IngredientRepository
}

func NewProductService(
	pr *repositories.ProductRepository,
	ir *repositories.IngredientRepository,
) *ProductService {
	return &ProductService{
		product_repo:    pr,
		ingredient_repo: ir,
	}
}

func (s *ProductService) GetPaginatedProducts(ctx context.Context, page, perPage int) (dto.PaginatedResponse, error) {
	products, totalCount, err := s.product_repo.GetProducts(ctx, page, perPage)
	if err != nil {
		return dto.PaginatedResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

	return dto.PaginatedResponse{
		Products:   products,
		TotalCount: totalCount,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *ProductService) GetProductByBarcode(ctx context.Context, barcode string) (dto.ProductDTO, error) {
	return s.product_repo.GetProductByBarcode(ctx, barcode)
}

func (s *ProductService) AnalyzeIngredients(ctx context.Context, aq dto.AnalyzeRequest) (dto.AnalyzedResponse, error) {
	ingerdients := s.ingredient_repo.GetIngredientsByIDs(aq)
	var totalRisk float64
	var analyzedIngredients []dto.AnalyzedIngredient
	for _, ingredient := range ingerdients {
		analyzedIngredient := dto.AnalyzedIngredient{
			Name:      ingredient.Name,
			RiskScore: ingredient.RiskScore,
			RiskType:  ingredient.RiskType,
		}
		analyzedIngredients = append(analyzedIngredients, analyzedIngredient)
		totalRisk += analyzedIngredient.RiskScore
	}

	overallRisk := totalRisk / float64(len(ingerdients))
	return dto.AnalyzedResponse{
		Ingredients: analyzedIngredients,
		OverallRisk: overallRisk,
	}, nil
}
