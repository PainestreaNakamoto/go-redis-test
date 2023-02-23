package main

import (
	"errors"
	"goredblue/repositories"
	"goredblue/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	db := initDatabase()
	app := fiber.New()
	redisClient := initRedis()
	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	app.Get("/hello", func(c *fiber.Ctx) error {

	data ,err := productService.GetProducts()
	if err != nil {
		return errors.New("errro")
	}
		return c.JSON(data)
	})
	app.Listen(":8000")

}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/infinitas")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
