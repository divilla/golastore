package category_list_breadcrumbs

import (
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	IData interface {
		WebPage() *view_model.WebPage
		CategoryList() *view_model.CategoryList
	}

	component struct {
		breadcrumbsHome string
		view_model.CategoryList
	}
)

func New(data IData) *html.Component {
	c := &component{
		data.WebPage().BreadcrumbsHome,
		*data.CategoryList(),
	}

	if len(c.CurrentCategory.Path) == 0 {
		return &html.Component{}
	}

	return c.html()
}

func (c *component) url(slug string) string {
	return c.URLBuilder(slug)
}

func (c *component) html() *html.Component {
	items := make([]html.Renderer, len(c.CurrentCategory.Path))
	for key, item := range c.CurrentCategory.Path {
		if key == 0 {
			items[key] =
				e.Li().Children(
					e.A(a.Href("/")).Text(c.breadcrumbsHome),
				)
		} else {
			items[key] =
				e.Li().Children(
					e.A(a.Href("/c/" + item.Slug)).Text(item.Name),
				)
		}
	}

	items = append(items,
		e.Li(a.Class("is-active")).Children(
			e.A(a.AriaCurrent("page")).Children(
				e.H1().Text(c.CurrentCategory.Name),
			),
		),
	)

	return html.NewComponent(
		e.Nav(a.Class("breadcrumb has-bullet-separator"), a.AriaLabel("breadcrumbs")).Children(
			e.Ul().Children(
				items...,
			),
		),
	)
}
