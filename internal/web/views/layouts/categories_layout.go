package layouts

import (
	"github.com/divilla/golastore/internal/domain"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	ICategoriesLayoutData interface {
		Title() string
		SelectedSlug() string
		ListedCategory() *domain.TaxonomyItem
	}
)

func NewCategoriesLayout(model ICategoriesLayoutData, view html.IView) *html.Layout {
	var path []html.Renderer
	var children []html.Renderer
	cat := model.ListedCategory()
	slug := model.SelectedSlug()
	chevronDown := "<i class=\"fas fa-chevron-down\" style=\"margin-right: 6px\"></i> "

	for k, v := range cat.Path {
		if k == len(cat.Path)-1 {
			break
		}
		path = append(path,
			e.Li().Children(
				e.A(a.Href("/c/"+v.Slug)).Children(
					e.Strong().Text(chevronDown+v.Name),
				),
			),
		)
	}

	for _, v := range cat.Children {
		children = append(children,
			e.Li().Children(
				e.A(a.Href("/c/"+v.Slug), a.Class(d.Ifs(v.Slug == slug, "is-active").String())).
					Text(v.Name),
			),
		)
	}

	if cat.Slug != slug {
		path = append(path,
			e.Li().Children(
				e.A(a.Href("/c/"+slug)).
					Text(chevronDown+cat.Name),
			),
			e.Li().Children(
				e.Ul(a.Class("menu-list")).Children(
					children...,
				),
			),
		)
	} else {
		path = append(path,
			e.Li().Children(
				e.A(a.Href("/c/"+slug), a.Class("is-active")).
					Text(chevronDown+cat.Name),
				e.Ul().Children(
					children...,
				),
			),
		)
	}

	return NewMainLayout(model, html.NewLayout(
		e.Div(a.Class("columns")).Children(
			e.Div(a.Class("column is-one-quarter"), a.Style("max-width: 300px")).Children(
				e.Aside(a.Class("menu")).Children(
					e.Ul(a.Class("menu-list")).Children(
						path...,
					),
				),
			),
			e.Div(a.Class("column")).Children(
				view,
			),
		),
	))
}