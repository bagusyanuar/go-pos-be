package admin

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/provider"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/middleware"
)

func RegisterRoutes(
	config *config.AppConfig,
	handlers *provider.Handlers,
) {
	app := config.App

	jwtMiddleware := middleware.VerifyJWT(config)
	privateApi := app.Group("/", jwtMiddleware)

	productCategories := privateApi.Group("/product-category")
	productCategories.Get("/", handlers.ProductCategory.Find)
	productCategories.Post("/", handlers.ProductCategory.Create)
	productCategories.Get("/:id", handlers.ProductCategory.FindByID)
}
