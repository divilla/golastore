package catalog_service

import (
	"context"
	"github.com/divilla/golastore/internal/cache"
	"github.com/divilla/golastore/internal/repository"
	"github.com/divilla/golastore/internal/view_model"
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
	category, err := s.taxCache.Get(dto.Category)
	if err != nil {
		return nil, err
	}
	categoryVM := view_model.NewCategoryVM(category)

	totalItems, err := s.productRepo.SearchCount(ctx, categoryVM.CurrentCategory.ShortId())
	if err != nil {
		return nil, err
	}
	pagination := view_model.NewPagination(dto.Page, totalItems)

	products, err := s.productRepo.Search(ctx, categoryVM.CurrentCategory.ShortId(), pagination.Limit(), pagination.Offset())
	if err != nil {
		return nil, err
	}

	return &CategoryModel{
		title:      s.appCache.Title(),
		Category:   categoryVM,
		Pagination: pagination,
		Products:   products,
	}, err
}
