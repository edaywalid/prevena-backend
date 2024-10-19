package repositories

import (
	"context"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/dto"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	product_collection    *mongo.Collection
	ingredient_collection *mongo.Collection
}

func NewProductRepository(
	product_collection *mongo.Collection,
	ingredient_collection *mongo.Collection,
) *ProductRepository {
	return &ProductRepository{
		product_collection:    product_collection,
		ingredient_collection: ingredient_collection,
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	_, err := r.product_collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) GetProducts(ctx context.Context, page, perPage int) ([]dto.ProductDTO, int64, error) {
	skip := int64((page - 1) * perPage)
	limit := int64(perPage)

	totalCount, err := r.product_collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := r.product_collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	var productsDTO []dto.ProductDTO

	for _, product := range products {
		var ingredients []models.Ingredient

		filter := bson.M{"_id": bson.M{"$in": product.Ingredients}}
		ingredientCursor, err := r.ingredient_collection.Find(ctx, filter)
		if err != nil {
			return nil, 0, err
		}
		if err = ingredientCursor.All(ctx, &ingredients); err != nil {
			return nil, 0, err
		}

		productDTO := dto.ProductDTO{
			Barcode:     product.Barcode,
			Name:        product.Name,
			Brand:       product.Brand,
			Ingredients: ingredients, OverallRisk: product.OverallRisk,
		}

		productsDTO = append(productsDTO, productDTO)
	}

	return productsDTO, totalCount, nil
}

func (r *ProductRepository) GetProductByBarcode(ctx context.Context, barcode string) (dto.ProductDTO, error) {
	var product models.Product
	err := r.product_collection.FindOne(ctx, bson.M{"_id": barcode}).Decode(&product)
	if err != nil {
		return dto.ProductDTO{}, err
	}

	var ingredients []models.Ingredient
	filter := bson.M{"_id": bson.M{"$in": product.Ingredients}}
	ingredientCursor, err := r.ingredient_collection.Find(ctx, filter)
	if err != nil {
		return dto.ProductDTO{}, err
	}
	if err = ingredientCursor.All(ctx, &ingredients); err != nil {
		return dto.ProductDTO{}, err
	}

	productDTO := dto.ProductDTO{
		Barcode:     product.Barcode,
		Name:        product.Name,
		Brand:       product.Brand,
		Ingredients: ingredients,
		OverallRisk: product.OverallRisk,
	}

	return productDTO, nil
}
