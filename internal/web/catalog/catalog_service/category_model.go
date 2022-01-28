package catalog_service

import "github.com/divilla/golastore/internal/domain"

type (
	CategoryDTO struct {
		Search       string `param:"search"`
		CategorySlug string `param:"category"`
	}

	CategoryModel struct {
		title    string
		category *domain.TaxonomyItem
		products string
	}
)

func (m *CategoryModel) Title() string {
	return m.title
}

func (m *CategoryModel) Category() *domain.TaxonomyItem {
	return m.category
}
