package products

import (
	"github.com/dacharat/wire-gin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type ProductRepositoryInterface interface {
// 	FindAll() ([]Product, error)
// 	FindByID(id string) (*Product, error)
// }

type ProductRepository struct {
	DB *mongo.Collection
}

func ProvideProductRepository(DB *mongo.Database) ProductRepository {
	return ProductRepository{DB: DB.Collection("wire")}
}

func (p *ProductRepository) FindAll() ([]Product, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	var products []Product
	cur, err := p.DB.Find(ctx, bson.M{})
	if err != nil {
		return products, err
	}

	err = cur.All(ctx, &products)

	return products, err
}

func (p *ProductRepository) FindByID(id string) (*Product, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product *Product
	err = p.DB.FindOne(ctx, bson.M{"_id": objectID}).Decode(product)

	return product, err
}

func (p *ProductRepository) Save(product *Product) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := p.DB.InsertOne(ctx, product)
	return err
}
