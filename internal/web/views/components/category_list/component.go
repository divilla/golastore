package category_list

import (
	"fmt"
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	IData interface {
		CategoryList() *view_model.CategoryList
	}

	component struct {
		view_model.CategoryList
	}
)

func New(data IData) *html.Component {
	c := &component{
		*data.CategoryList(),
	}

	return c.html()
}

func (c *component) url(slug string) string {
	return c.URLBuilder(slug)
}

func (c *component) html() *html.Component {
	chevronDown := "<i class=\"fas fa-chevron-down\" style=\"margin-right: 6px\"></i> "

	path := make([]html.Renderer, len(c.ListedCategory.Path))
	for k, v := range c.ListedCategory.Path {
		path[k] =
			e.Li().Children(
				e.A(a.Href(c.URLBuilder(v.Slug))).Children(
					e.Strong().Text(chevronDown + v.Name),
				),
			)
	}

	items := make([]html.Renderer, len(c.ListedCategory.Children))
	for k, v := range c.ListedCategory.Children {
		if v.Slug == c.CurrentCategorySlug {
			items[k] =
				e.Li().Children(
					e.A(a.Href(c.URLBuilder(v.Slug)), a.Class("is-active")).
						Text(fmt.Sprintf("%s (%s)", v.Name, v.TotalProducts())),
				)
		} else {
			items[k] =
				e.Li().Children(
					e.A(a.Href(c.URLBuilder(v.Slug))).
						Text(fmt.Sprintf("%s (%s)", v.Name, v.TotalProducts())),
				)
		}
	}

	if c.ListedCategory.Slug != c.CurrentCategorySlug {
		path = append(path,
			e.Li().Children(
				e.A(a.Href(c.URLBuilder(c.ListedCategory.Slug))).
					Text(chevronDown+c.ListedCategory.Name),
			),
			e.Li().Children(
				e.Ul(a.Class("menu-list")).Children(
					items...,
				),
			),
		)
	} else {
		path = append(path,
			e.Li().Children(
				e.A(a.Href(c.URLBuilder(c.ListedCategory.Slug)), a.Class("is-active")).
					Text(chevronDown+c.ListedCategory.Name),
				e.Ul().Children(
					items...,
				),
			),
		)
	}

	return html.NewComponent(
		e.Aside(a.Class("menu")).Children(
			e.Ul(a.Class("menu-list")).Children(
				path...,
			),
		),
	)
}
