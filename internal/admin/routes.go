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
	materialCategories.Put("/:id", handlers.MaterialCategory.Update)
	materialCategories.Delete("/:id", handlers.MaterialCategory.Delete)

	productCategories := privateApi.Group("/product-category")
	productCategories.Get("/", handlers.ProductCategory.Find)
	productCategories.Post("/", handlers.ProductCategory.Create)
	productCategories.Get("/:id", handlers.ProductCategory.FindByID)
	productCategories.Put("/:id", handlers.ProductCategory.Update)
	productCategories.Delete("/:id", handlers.ProductCategory.Delete)

	unit := privateApi.Group("/unit")
	unit.Get("/", handlers.Unit.Find)
	unit.Post("/", handlers.Unit.Create)
	unit.Get("/:id", handlers.Unit.FindByID)
	unit.Put("/:id", handlers.Unit.Update)
	unit.Delete("/:id", handlers.Unit.Delete)

	material := privateApi.Group("/material")
	material.Get("/", handlers.Material.Find)
	material.Post("/", handlers.Material.Create)
	material.Get("/:id", handlers.Material.FindByID)
	material.Put("/:id", handlers.Material.Update)
	material.Delete("/:id", handlers.Material.Delete)
	material.Put("/:id/image", handlers.Material.UploadImage)
	material.Put("/:id/unit", handlers.Material.AppendUnit)

	materialInventory := privateApi.Group("/material-inventory")
	materialInventory.Get("/", handlers.MaterialInventory.Find)

}
