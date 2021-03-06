package routers

import (
	"github.com/dacharat/wire-gin/products"
	"github.com/gin-gonic/gin"
)

func New(h *products.ProductAPI) *gin.Engine {
	router := gin.New()

	p := router.Group("/api/v1/products")
	{
		p.GET("", h.FindAll)
		p.GET("/:id", h.FindByID)

		p.POST("", h.Create)
	}

	return router
}
