package dto

import "github.com/edaywalid/pinktober-hackathon-backend/internal/models"

type ProductDTO struct {
	Barcode     string              `json:"barcode" bson:"_id"`
	Name        string              `json:"name" bson:"name"`
	Brand       string              `json:"brand" bson:"brand"`
	Ingredients []models.Ingredient `json:"ingredients" bson:"ingredients"`
	OverallRisk float64             `json:"overall_risk" bson:"overall_risk"`
}
