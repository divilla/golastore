package catalog

import (
	"github.com/divilla/golastore/framework/middleware"
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/labstack/echo/v4"
)

type (
	catalogController struct {
		service *catalog_service.CatalogService
	}
)

func NewController(e *echo.Echo, service *catalog_service.CatalogService) {
	c := &catalogController{
		service: service,
	}

	e.GET("/", middleware.UseCustomContext(c.index))

	group := e.Group("/catalog")
	group.GET("/", middleware.UseCustomContext(c.index))
}

func (c *catalogController) index(ctx *middleware.CustomContext) error {
	return ctx.NoContent(200)
}
