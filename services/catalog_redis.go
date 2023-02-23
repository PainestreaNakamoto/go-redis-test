package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredblue/repositories"
	"time"
	"github.com/go-redis/redis/v8"
)


type catalogServiceRedis struct {
	repo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(repo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{repo: repo, redisClient: redisClient }
}

func (self catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "services::GetProducts"

	// Redis 
	product_json, err:= self.redisClient.Get(context.Background(), key).Result()
	fmt.Println(err)
	if err == nil {
		if json.Unmarshal([]byte(product_json),&products) == nil{
			fmt.Println("redis")
			return products, nil
		}
	}

	product_db ,  err := self.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _ , p := range product_db {
		products = append(products, Product{p.ID,p.Name,p.Quantity})
	}

	if data, err := json.Marshal(products); err == nil {
		self.redisClient.Set(context.Background(), key, string(data), time.Second * 10)
	}

	fmt.Println("DB")

	return products, nil
}


