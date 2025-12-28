package admin

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/provider"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func RegisterRoutes(
	config *config.AppConfig,
	handlers *provider.Handlers,
) {
	app := config.App

	apiV1 := app.Group("/v1")

	adminApi := apiV1.Group("/admin")

	productCategories := adminApi.Group("/product-category")
	productCategories.Get("/", handlers.ProductCategory.FindAll)
	productCategories.Post("/", handlers.ProductCategory.Create)
	productCategories.Get("/:id", handlers.ProductCategory.FindByID)
}
