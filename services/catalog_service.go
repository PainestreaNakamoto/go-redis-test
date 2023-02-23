package services

import (
	"goredblue/repositories"
)

type catalogService struct {
	repo repositories.ProductRepository
}


func NewCatalogService(product_repo repositories.ProductRepository) CatalogService{
	return catalogService{repo: product_repo}
}

func (self catalogService) GetProducts() (products []Product,err error) {
	product_db, err := self.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _ , p := range product_db {
		products = append(products, Product{ID: p.ID,Name: p.Name, Quantity: p.Quantity})
	}
	return products,nil
}

