package products

import (
	"github.com/dacharat/wire-gin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(id string) (*Product, error)
	Save(product *Product) error
}

type ProductRepository struct {
	db *mongo.Collection
}

func ProvideProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{db: db.Collection("wire")}
}

func (p *ProductRepository) FindAll() ([]Product, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	var products []Product
	cur, err := p.db.Find(ctx, bson.M{})
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
	err = p.db.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)

	return product, err
}

func (p *ProductRepository) Save(product *Product) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	product.ID = primitive.NewObjectID()

	_, err := p.db.InsertOne(ctx, product)
	return err
}
