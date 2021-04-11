package products

type ProductService struct {
	ProductRepository ProductRepository
}

func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
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
