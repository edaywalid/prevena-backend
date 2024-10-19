package seed

import (
	"context"
	"log"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/models"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/repositories"
)

type Seeder struct {
	productRepo    *repositories.ProductRepository
	ingredientRepo *repositories.IngredientRepository
}

func NewSeeder(productRepo *repositories.ProductRepository, ingredientRepo *repositories.IngredientRepository) *Seeder {
	return &Seeder{
		productRepo:    productRepo,
		ingredientRepo: ingredientRepo,
	}
}

func (s *Seeder) SeedAll(ctx context.Context) error {
	if err := s.SeedIngredients(ctx); err != nil {
		return err
	}
	if err := s.SeedProducts(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Seeder) SeedIngredients(ctx context.Context) error {
	var ingredients = []models.Ingredient{
		{ID: "ing001", Name: "Water", RiskScore: 0.0, RiskType: "Safe"},
		{ID: "ing002", Name: "Sodium Laureth Sulfate", RiskScore: 0.3, RiskType: "Low"},
		{ID: "ing003", Name: "Cocamidopropyl Betaine", RiskScore: 0.2, RiskType: "Low"},
		{ID: "ing004", Name: "Glycerin", RiskScore: 0.1, RiskType: "Safe"},
		{ID: "ing005", Name: "Fragrance", RiskScore: 0.7, RiskType: "Moderate"},
		{ID: "ing006", Name: "Citric Acid", RiskScore: 0.1, RiskType: "Safe"},
		{ID: "ing007", Name: "Methylparaben", RiskScore: 0.6, RiskType: "Moderate"},
		{ID: "ing008", Name: "Propylparaben", RiskScore: 0.6, RiskType: "Moderate"},
		{ID: "ing009", Name: "Salicylic Acid", RiskScore: 0.4, RiskType: "Low"},
		{ID: "ing010", Name: "Retinol", RiskScore: 0.5, RiskType: "Moderate"},
	}

	for _, ingredient := range ingredients {
		if err := s.ingredientRepo.CreateIngredient(ctx, &ingredient); err != nil {
			log.Printf("Error seeding ingredient %s: %v", ingredient.ID, err)
			return err
		}
	}

	log.Println("Ingredients seeded successfully")
	return nil
}

func (s *Seeder) SeedProducts(ctx context.Context) error {
	var products = []models.Product{
		{
			Barcode: "123456789",
			Name:    "Gentle Cleansing Shampoo",
			Brand:   "NatureCare",
			Ingredients: []string{
				"ing001", "ing002", "ing003", "ing004", "ing005", "ing006",
			},
			OverallRisk: 0.23,
		},
		{
			Barcode: "987654321",
			Name:    "Moisturizing Face Cream",
			Brand:   "LuxeSkin",
			Ingredients: []string{
				"ing001", "ing004", "ing005", "ing007", "ing008",
			},
			OverallRisk: 0.40,
		},
		{
			Barcode: "456789123",
			Name:    "Acne Treatment Gel",
			Brand:   "ClearSkin",
			Ingredients: []string{
				"ing001", "ing004", "ing009", "ing006",
			},
			OverallRisk: 0.15,
		},
		{
			Barcode: "789123456",
			Name:    "Anti-Aging Serum",
			Brand:   "TimelessBeauty",
			Ingredients: []string{
				"ing001", "ing004", "ing010", "ing005", "ing006",
			},
			OverallRisk: 0.36,
		},
	}

	for _, product := range products {
		if err := s.productRepo.CreateProduct(ctx, &product); err != nil {
			log.Printf("Error seeding product %s: %v", product.Barcode, err)
			return err
		}
	}

	log.Println("Products seeded successfully")
	return nil
}
