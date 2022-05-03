package layouts

import (
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/internal/web/views/components/category_list"
	"github.com/divilla/golastore/internal/web/views/components/category_list_breadcrumbs"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	ICategoriesLayoutData interface {
		WebPage() *view_model.WebPage
		CategoryList() *view_model.CategoryList
		BreadCrumbs() *view_model.Breadcrumbs
	}
)

func NewCategoriesLayout(data ICategoriesLayoutData, view html.IView) *html.Layout {
	return NewMainLayout(data,
		html.NewLayout(
			e.Div(a.Class("columns")).Children(
				e.Div(a.Class("column cols-xl-6 cols-lg-5 cols-md-4 cols-sm-3 cols-xs-2 cols-1")).Children(
					category_list.New(data),
				),
				e.Div(a.Class("column")).Children(
					category_list_breadcrumbs.New(data),
					view,
				),
			),
			e.Script(a.Type("module"), a.Src("/assets/js/category.js")),
		),
	)
}
