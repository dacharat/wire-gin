package products

import "github.com/google/wire"

var SetService = wire.NewSet(ProvideProductService)

type ProviderProductService interface {
	FindAll() ([]Product, error)
	FindByID(id string) (*Product, error)
	Save(product *Product) error
}
type ProductService struct {
	ProductRepository ProviderProductRepository
}

func ProvideProductService(p ProviderProductRepository) *ProductService {
	return &ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() ([]Product, error) {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id string) (*Product, error) {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product *Product) error {

	return p.ProductRepository.Save(product)
}
