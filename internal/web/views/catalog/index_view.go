package catalog

import (
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/d"
)

func NewIndexView(model *catalog_service.CategoryModel) html.IView {
	return layouts.NewCategoriesLayout(model, html.NewView(d.Block(model.Title())))
}
