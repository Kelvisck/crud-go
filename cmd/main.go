package main

import (
	router "go-api"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	/*dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}*/

	//ProductRepository := repository.NewProductRepository(dbConnection)
	client, err := db.ConnectMongo()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(client)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//server.GET("/products", ProductController.GetProducts)
	//server.POST("/products", ProductController.CreateProduct)

	r := router.SetupRoutes(&ProductController)
	r.Run(":8000")

	//server.Run(":8000")

}
