package catalog_service

import (
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/divilla/golastore/internal/view_model"
)

type (
	CategoryDTO struct {
		Search   string `param:"search"`
		Category string `param:"category"`
		Page     int64  `param:"page"`
	}

	CategoryModel struct {
		title      string
		Category   *view_model.CategoryVM
		Pagination *view_model.PaginationVM
		Products   []*domain_model.ListProduct
	}
)

func (m *CategoryModel) Title() string {
	return m.title
}

func (m *CategoryModel) CategorySlug() string {
	return m.Category.CurrentCategorySlug
}

func (m *CategoryModel) CurrentCategory() *domain_model.TaxonomyItem {
	return m.Category.CurrentCategory
}

func (m *CategoryModel) ListedCategory() *domain_model.TaxonomyItem {
	return m.Category.ListedCategory
}
