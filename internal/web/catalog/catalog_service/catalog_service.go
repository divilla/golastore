package catalog_service

import (
	"github.com/divilla/golastore/internal/cache"
	"github.com/google/uuid"
)

type (
	CatalogService struct {
		appCache *cache.App
		taxCache *cache.Taxonomy
	}

	ProductItem struct {
		Id   uuid.UUID
		Name string
	}
)

func NewCatalogService(appCache *cache.App, taxCache *cache.Taxonomy) *CatalogService {
	return &CatalogService{
		appCache: appCache,
		taxCache: taxCache,
	}
}

func (s *CatalogService) Category(c *CategoryDTO) *CategoryModel {
	m := &CategoryModel{}

	m.title = s.appCache.Title()
	if c.CategorySlug == "" || c.CategorySlug == s.taxCache.ProductCategoriesRoot().Slug {
		m.listedCategory = s.taxCache.ProductCategoriesRoot()
		m.selectedSlug = m.listedCategory.Slug
	} else {
		m.selectedSlug = c.CategorySlug
		m.listedCategory = s.taxCache.Get(m.selectedSlug)
		m.title = m.listedCategory.Name + " - " + m.title
	}
	if len(m.listedCategory.Children) == 0 && m.listedCategory.ParentSlug.String != "" {
		m.listedCategory = s.taxCache.Get(m.listedCategory.ParentSlug.String)
	}

	return m
}
