package maintenance_controller

import (
	"github.com/divilla/golastore/framework/middleware"
	"github.com/divilla/golastore/internal/web/maintenance/maintenance_service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	maintenanceController struct {
		service *maintenance_service.MaintenanceService
	}
)

func New(echo *echo.Echo, service *maintenance_service.MaintenanceService) {
	c := &maintenanceController{
		service: service,
	}

	group := echo.Group("/maintenance")
	group.GET("/rebuild-taxonomy-slugs", middleware.UseCustomContext(c.rebuildTaxonomySlugs))
	group.GET("/rebuild-other-taxonomy-slugs", middleware.UseCustomContext(c.rebuildOtherTaxonomySlugs))
	group.GET("/rebuild-taxonomy-parents", middleware.UseCustomContext(c.rebuildTaxonomyParents))
	group.GET("/fix-products", middleware.UseCustomContext(c.fixProducts))
}

func (c *maintenanceController) rebuildTaxonomySlugs(ctx *middleware.CustomContext) error {
	if err := c.service.RebuildTaxonomySlugs(ctx.Request().Context()); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *maintenanceController) rebuildOtherTaxonomySlugs(ctx *middleware.CustomContext) error {
	if err := c.service.RebuildOtherTaxonomySlugs(ctx.Request().Context()); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *maintenanceController) rebuildTaxonomyParents(ctx *middleware.CustomContext) error {
	if err := c.service.RebuildTaxonomyParents(ctx.Request().Context()); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *maintenanceController) fixProducts(ctx *middleware.CustomContext) error {
	if err := c.service.FixProducts(ctx.Request().Context()); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
