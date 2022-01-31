package catalog_service

import "github.com/divilla/golastore/internal/domain"

type (
	CategoryDTO struct {
		Search       string `param:"search"`
		CategorySlug string `param:"category"`
	}

	CategoryModel struct {
		title            string
		selectedSlug     string
		selectedCategory *domain.TaxonomyItem
		listedCategory   *domain.TaxonomyItem
		productsList     []*domain.ListProduct
	}
)

func (m *CategoryModel) Title() string {
	return m.title
}

func (m *CategoryModel) SelectedSlug() string {
	return m.selectedSlug
}

func (m *CategoryModel) SelectedCategory() *domain.TaxonomyItem {
	return m.selectedCategory
}

func (m *CategoryModel) ListedCategory() *domain.TaxonomyItem {
	return m.listedCategory
}

func (m *CategoryModel) ProductsList() []*domain.ListProduct {
	return m.productsList
}
