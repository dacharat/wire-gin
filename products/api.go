package products

import (
	"net/http"

	"github.com/dacharat/wire-gin/utils"
	"github.com/gin-gonic/gin"
)

type ProviderProductAPI interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Create(c *gin.Context)
}

type ProductAPI struct {
	ProductService ProviderProductService
}

func ProvideProductAPI(p ProviderProductService) *ProductAPI {
	return &ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products, err := p.ProductService.FindAll()

	if err != nil {
		utils.Handler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id := c.Param("id")
	product, err := p.ProductService.FindByID(id)
	if err != nil {
		utils.Handler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(*product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		utils.Handler(c, err)
		return
	}

	product := ToProduct(productDTO)
	err = p.ProductService.Save(&product)
	if err != nil {
		utils.Handler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}
