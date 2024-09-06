package routers

import (
	"product-ms/apps/controllers/handlers"
	"product-ms/apps/models/repositories"
	"product-ms/apps/models/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *fiber.App {
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewHandler(productService)

	r := fiber.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	r.Post("/api/v1/product/create", productHandler.CreateProduct)
	r.Get("/api/v1/product/:id", productHandler.GetProductByID)
	r.Get("/api/v1/products", productHandler.GetProducts)
	r.Put("/api/v1/product/:id", productHandler.UpdateProduct)
	r.Delete("/api/v1/product/:id", productHandler.DeleteProduct)

	return r
}
