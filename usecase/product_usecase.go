package usecase

import (
	"go-api/model"
	"go-api/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUsecase struct {
	repository *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	return pu.repository.CreateProduct(product)

}

func (pu *ProductUsecase) GetProductByID(id primitive.ObjectID) (model.Product, error) {
	return pu.repository.GetProductByID(id)
}

func (pu *ProductUsecase) UpdateProduct(id primitive.ObjectID, product model.Product) (model.Product, error) {
	return pu.repository.UpdateProduct(id, product)
}

func (pu *ProductUsecase) DeleteProduct(id primitive.ObjectID) error {
	return pu.repository.DeleteProduct(id)
}
