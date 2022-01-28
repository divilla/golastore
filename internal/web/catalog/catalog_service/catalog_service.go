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
	m := &CategoryModel{
		title: s.appCache.Title(),
	}

	if c.CategorySlug == "" {
		m.category = s.taxCache.ProductCategories()
	} else {
		m.category = s.taxCache.Get(c.CategorySlug)
	}

	return m
}
