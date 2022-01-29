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
		Category() *domain.TaxonomyItem
	}
)

func NewCategoriesLayout(data ICategoriesLayoutData, view html.IView) *html.Layout {
	var parents []html.Renderer
	var children []html.Renderer
	item := data.Category()
	selectedItem := data.Category()
	chevronDown := "<i class=\"fas fa-chevron-down\" style=\"margin-right: 6px\"></i> "

	if len(item.Children) == 0 {
		l := len(item.Parents)
		if l > 0 {
			item = item.Parents[l-1]
		}
	}

	for _, v := range item.Parents {
		parents = append(parents,
			e.Li().Children(
				e.A(a.Href("/c/"+v.Slug)).Children(e.Strong().Text(chevronDown+v.Name)),
			),
		)
	}

	for _, v := range item.Children {
		children = append(children,
			e.Li().Children(
				e.A(a.Href("/c/"+v.Slug), a.Class(d.Ifs(v.Slug == selectedItem.Slug, "is-active").String())).
					Text(v.Name),
			),
		)
	}

	slug := item.Slug
	if item.ParentSlug.String != "" {
		slug = item.ParentSlug.String
	}
	if item.Slug != selectedItem.Slug {
		parents = append(parents,
			e.Li().Children(
				e.A(a.Href("/c/"+slug)).
					Text(chevronDown+item.Name),
			),
			e.Li().Children(
				e.Ul(a.Class("menu-list")).Children(
					children...,
				),
			),
		)
	} else {
		parents = append(parents,
			e.Li().Children(
				e.A(a.Href("/c/"+slug), a.Class("is-active")).
					Text(chevronDown+item.Name),
				e.Ul().Children(
					children...,
				),
			),
		)
	}

	return NewMainLayout(data, html.NewLayout(
		e.Div(a.Class("columns")).Children(
			e.Div(a.Class("column is-one-quarter"), a.Style("max-width: 300px")).Children(
				e.Aside(a.Class("menu")).Children(
					e.Ul(a.Class("menu-list")).Children(
						parents...,
					),
				),
			),
			e.Div(a.Class("column")).Children(
				view,
			),
		),
	))
}
