package layouts

import (
	"github.com/divilla/golastore/internal/domain"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	ICategoriesLayoutData interface {
		Title() string
		Category() *domain.TaxonomyItem
	}
)

func NewCategoriesLayout(data ICategoriesLayoutData, view html.IView) *html.Layout {
	var categories []html.Renderer
	for _, v := range data.Category().Children {
		categories = append(categories,
			e.Li().Children(
				e.A(a.Href("/c/"+v.Slug)).Text(v.Name),
			),
		)
	}

	return NewMainLayout(data, html.NewLayout(
		e.Div(a.Class("columns")).Children(
			e.Div(a.Class("column is-one-quarter")).Children(
				e.Aside(a.Class("menu")).Children(
					e.P(a.Class("menu-label")).
						Text(data.Category().Name),
					e.Ul(a.Class("menu-list")).Children(
						categories...,
					),
				),
			),
			e.Div(a.Class("column is-three-quarters")).Children(
				view,
			),
		),
	))
}
