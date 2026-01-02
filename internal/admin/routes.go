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

	materialCategories := privateApi.Group("/material-category")
	materialCategories.Get("/", handlers.MaterialCategory.Find)
	materialCategories.Post("/", handlers.MaterialCategory.Create)
	materialCategories.Get("/:id", handlers.MaterialCategory.FindByID)

	productCategories := privateApi.Group("/product-category")
	productCategories.Get("/", handlers.ProductCategory.Find)
	productCategories.Post("/", handlers.ProductCategory.Create)
	productCategories.Get("/:id", handlers.ProductCategory.FindByID)
}
