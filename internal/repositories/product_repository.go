package repositories

import (
	"context"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) *ProductRepository {
	return &ProductRepository{collection: collection}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) GetProducts(ctx context.Context, page, perPage int) ([]models.Product, int64, error) {
	skip := int64((page - 1) * perPage)
	limit := int64(perPage)

	totalCount, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

func (r *ProductRepository) GetProductByBarcode(ctx context.Context, barcode string) (models.Product, error) {
	var product models.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": barcode}).Decode(&product)
	return product, err
}
