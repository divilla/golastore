package catalog_service

import "github.com/divilla/golastore/internal/domain"

type (
	CategoryDTO struct {
		Search       string `param:"search"`
		CategorySlug string `param:"category"`
		Page         int64  `param:"page"`
	}

	CategoryModel struct {
		title            string
		selectedSlug     string
		selectedCategory *domain.TaxonomyItem
		listedCategory   *domain.TaxonomyItem
		currentPage      int64
		productsPerPage  int64
		totalProducts    int64
		totalPages       int64
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

func (m *CategoryModel) CurrentPage() int64 {
	return m.currentPage
}

func (m *CategoryModel) ItemsPerPage() int64 {
	return m.productsPerPage
}

func (m *CategoryModel) TotalPages() int64 {
	return m.totalPages
}

func (m *CategoryModel) TotalItems() int64 {
	return m.totalProducts
}

func (m *CategoryModel) ProductsList() []*domain.ListProduct {
	return m.productsList
}
