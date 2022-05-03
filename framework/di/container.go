package di

import (
	"github.com/divilla/golastore/framework/config"
	"github.com/divilla/golastore/internal/cache"
	"github.com/divilla/golastore/internal/repository"
	"github.com/divilla/golastore/internal/web/cart"
	"github.com/divilla/golastore/internal/web/catalog"
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/maintenance/maintenance_controller"
	"github.com/divilla/golastore/internal/web/maintenance/maintenance_service"
	"github.com/divilla/golastore/pkg/logger"
	"github.com/divilla/golastore/pkg/postgres"
	"github.com/labstack/echo/v4"
)

type (
	Container struct {
		logger *logger.Logger
		config *config.Config
		pool   *postgres.Pool
	}
)

func NewContainer(e *echo.Echo) *Container {
	log := logger.New()
	conf := config.New(log)
	pool := postgres.NewPool(conf.Dsn)

	appCache := cache.NewAppCache()

	taxonomyRepository := repository.NewTaxonomyRepository(pool)
	productRepository := repository.NewProductRepository(pool)
	//taxonomyService := service.NewTaxonomyService(taxonomyRepository)
	taxonomyCache := cache.NewTaxonomyCache(taxonomyRepository)

	catalogService := catalog_service.NewCatalogService(appCache, taxonomyCache, productRepository)
	catalog.NewController(e, catalogService)

	cartService := cart.NewService()
	cart.NewController(e, cartService)

	maintenanceService := maintenance_service.New(pool)
	maintenance_controller.New(e, maintenanceService)

	return &Container{
		logger: log,
		config: conf,
		pool:   pool,
	}
}

func (m *Container) Logger() *logger.Logger {
	return m.logger
}

func (m *Container) Config() *config.Config {
	return m.config
}

func (m *Container) Close() {
	m.logger.Close()
	m.pool.Close()
}
