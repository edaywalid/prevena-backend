package repositories

import (
	"context"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/dto"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IngredientRepository struct {
	collection *mongo.Collection
}

func NewIngredientRepository(collection *mongo.Collection) *IngredientRepository {
	return &IngredientRepository{collection: collection}
}
func (ir *IngredientRepository) GetIngredientsByIDs(ar dto.AnalyzeRequest) []models.Ingredient {
	var ingredients []models.Ingredient

	ids := make([]interface{}, len(ar.Ingerdients))

	for i, id := range ar.Ingerdients {
		ids[i] = id
	}

	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, _ := ir.collection.Find(nil, filter)
	cursor.All(nil, &ingredients)

	return ingredients
}

func (ir *IngredientRepository) CreateIngredient(ctx context.Context, ingredient *models.Ingredient) error {
	_, err := ir.collection.InsertOne(ctx, ingredient)
	return err
}
