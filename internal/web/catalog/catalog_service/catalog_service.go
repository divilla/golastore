package catalog_service

import (
	"github.com/divilla/golastore/internal/cache"
	"github.com/divilla/golastore/internal/domain"
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

	CategoryDTO struct {
		Search       string `json:"s"`
		CategorySlug string `json:"c"`
	}

	CategoryModel struct {
		Title    string
		Category *domain.TaxonomyItem
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
		Title: s.appCache.Title(),
	}

	if c.CategorySlug == "" {
		m.Category = s.taxCache.ProductCategories()
	}

	return m
}
