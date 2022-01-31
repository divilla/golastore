package catalog_service

import (
	"context"
	"github.com/divilla/golastore/internal/cache"
	"github.com/divilla/golastore/internal/repository"
	"strings"
)

type (
	CatalogService struct {
		appCache    *cache.App
		taxCache    *cache.Taxonomy
		productRepo *repository.Product
	}
)

func NewCatalogService(appCache *cache.App, taxCache *cache.Taxonomy, productRepo *repository.Product) *CatalogService {
	return &CatalogService{
		appCache:    appCache,
		taxCache:    taxCache,
		productRepo: productRepo,
	}
}

func (s *CatalogService) Category(ctx context.Context, dto *CategoryDTO) (*CategoryModel, error) {
	model := &CategoryModel{
		productsPerPage: 30,
	}

	model.title = s.appCache.Title()
	if dto.CategorySlug == "" || dto.CategorySlug == s.taxCache.ProductCategoriesRoot().Slug {
		model.selectedCategory = s.taxCache.ProductCategoriesRoot()
		model.listedCategory = s.taxCache.ProductCategoriesRoot()
		model.selectedSlug = model.listedCategory.Slug
	} else {
		model.selectedSlug = dto.CategorySlug
		model.selectedCategory = s.taxCache.Get(model.selectedSlug)
		model.listedCategory = s.taxCache.Get(model.selectedSlug)
		model.title = model.listedCategory.Name + " - " + model.title
	}
	if len(model.listedCategory.Children) == 0 && model.listedCategory.ParentSlug.String != "" {
		model.listedCategory = s.taxCache.Get(model.listedCategory.ParentSlug.String)
	}

	model.currentPage = dto.Page
	if model.currentPage < 1 {
		model.currentPage = 1
	}
	search := strings.ReplaceAll(model.SelectedCategory().Id.String(), "-", "")
	pl, pages, err := s.productRepo.Search(ctx, search, model.productsPerPage, model.currentPage)
	if err != nil {
		return nil, err
	}
	model.totalPages = pages
	model.productsList = pl

	return model, nil
}
