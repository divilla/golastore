package catalog_service

import "github.com/divilla/golastore/internal/domain"

type (
	CategoryDTO struct {
		Search       string `param:"search"`
		CategorySlug string `param:"category"`
	}

	CategoryModel struct {
		title          string
		selectedSlug   string
		listedCategory *domain.TaxonomyItem
		products       string
	}
)

func (m *CategoryModel) Title() string {
	return m.title
}

func (m *CategoryModel) SelectedSlug() string {
	return m.selectedSlug
}

func (m *CategoryModel) ListedCategory() *domain.TaxonomyItem {
	return m.listedCategory
}
