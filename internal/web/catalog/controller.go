package catalog

import (
	"github.com/divilla/golastore/framework/middleware"
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/catalog"
	"github.com/labstack/echo/v4"
	"net/http"
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
	e.GET("/c/:category", middleware.UseCustomContext(c.category))

	group := e.Group("/catalog")
	group.GET("/", middleware.UseCustomContext(c.index))
}

func (c *catalogController) index(ctx *middleware.CustomContext) error {
	var dto catalog_service.CategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return err
	}

	model, err := c.service.Category(ctx.Request().Context(), &dto)
	if err != nil {
		return err
	}

	return ctx.RenderView(http.StatusOK, catalog.NewIndexView(model))
}

func (c *catalogController) category(ctx *middleware.CustomContext) error {
	var dto catalog_service.CategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return err
	}

	model, err := c.service.Category(ctx.Request().Context(), &dto)
	if err != nil {
		return err
	}

	return ctx.RenderView(http.StatusOK, catalog.NewIndexView(model))
}
