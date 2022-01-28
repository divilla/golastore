package catalog

import (
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/d"
)

type (
	IndexView struct {
		html.View
		model *catalog_service.ProductItem
	}
)

func NewIndexView(layout *layouts.MainLayout) html.Renderer {
	return html.NewView(layout, d.Block("test"))
}
