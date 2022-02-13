package catalog_service

import (
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/divilla/golastore/internal/view_model"
)

type (
	CatalogCategoryDTO struct {
		Search       string `param:"search"`
		Category     string `param:"category"`
		Page         int64  `param:"page"`
		ItemsPerPage int64  `query:"itemsPerPage"`
	}

	CatalogCategoryModel struct {
		webPage      *view_model.WebPage
		categoryList *view_model.CategoryList
		breadcrumbs  *view_model.Breadcrumbs
		pagination   *view_model.Pagination
		listProducts []*domain_model.ListProduct
	}
)

func NewCatalogCategoryModel(
	webPage *view_model.WebPage,
	categoryList *view_model.CategoryList,
	breadcrumbs *view_model.Breadcrumbs,
	pagination *view_model.Pagination,
	listProducts []*domain_model.ListProduct,
) *CatalogCategoryModel {
	return &CatalogCategoryModel{
		webPage:      webPage,
		categoryList: categoryList,
		breadcrumbs:  breadcrumbs,
		pagination:   pagination,
		listProducts: listProducts,
	}
}

func (m *CatalogCategoryModel) WebPage() *view_model.WebPage {
	return m.webPage
}

func (m *CatalogCategoryModel) CategoryList() *view_model.CategoryList {
	return m.categoryList
}

func (m *CatalogCategoryModel) BreadCrumbs() *view_model.Breadcrumbs {
	return m.breadcrumbs
}

func (m *CatalogCategoryModel) Pagination() *view_model.Pagination {
	return m.pagination
}

func (m *CatalogCategoryModel) ListProducts() []*domain_model.ListProduct {
	return m.listProducts
}
