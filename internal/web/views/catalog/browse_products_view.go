package catalog

import (
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/components"
	"github.com/divilla/golastore/internal/web/views/components/pagination"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

func NewBrowseProductsView(m *catalog_service.CatalogCategoryModel) *html.Layout {
	paginationComponent := pagination.New(m)
	items := components.NewProductList(m.ListProducts())

	return layouts.NewCategoriesLayout(m,
		html.NewView(
			paginationComponent,
			e.Div(a.Class("columns items is-multiline")).Children(
				items...,
			),
			paginationComponent,
		),
	)
}
