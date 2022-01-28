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
	group.GET("/rebuild-taxonomy-parents", middleware.UseCustomContext(c.rebuildTaxonomyParents))
}

func (c *maintenanceController) rebuildTaxonomySlugs(ctx *middleware.CustomContext) error {
	if err := c.service.RebuildTaxonomySlugs(ctx.Request().Context()); err != nil {
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
