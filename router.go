package router

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(productController *controller.ProductController) *gin.Engine {
	r := gin.Default()

	// Grupo de rotas /products
	products := r.Group("/products")
	{
		products.GET("/", productController.GetProducts)
		products.POST("/", productController.CreateProduct)
		products.GET("/:id", productController.GetProductByID)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
	}

	return r
}
