package catalog_service

import (
	"context"
	"github.com/divilla/golastore/internal/cache"
	"github.com/divilla/golastore/internal/repository"
	"github.com/divilla/golastore/internal/view_model"
	"strconv"
)

const (
	rootProductCategory string = "product-categories-root"
	itemsPerPageDefault int64  = 30
	linkSpreadDefault   int64  = 4
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

func (s *CatalogService) CategoryProductList(ctx context.Context, dto *CatalogCategoryDTO) (*CatalogCategoryModel, error) {
	if dto.Page < 1 {
		dto.Page = 1
	}
	if dto.Category == "" {
		dto.Category = rootProductCategory
	}
	if dto.ItemsPerPage == 0 {
		dto.ItemsPerPage = itemsPerPageDefault
	}

	categoryURLBuilder := func(slug string) string {
		return "/c/" + slug + "/1"
	}
	category, err := s.taxCache.Get(dto.Category)
	if err != nil {
		return nil, err
	}
	categoryList := view_model.NewCategoryList(category, categoryURLBuilder)

	title := s.appCache.Title()
	if !category.Root {
		title = category.Name + " - " + title
	}
	webPage := view_model.NewWebPage(title, title)

	paginationURLBuilder := func(page int64) string {
		p := strconv.FormatInt(page, 10)
		return "/c/" + dto.Category + "/" + p
	}

	totalItems, err := s.productRepo.SearchCount(ctx, category.ShortId())
	if err != nil {
		return nil, err
	}
	pagination := view_model.NewPagination(dto.Page, dto.ItemsPerPage, totalItems, linkSpreadDefault, paginationURLBuilder)

	breadcrumbs := view_model.NewBreadcrumbsViewModel("Home")
	for _, cat := range category.Path {
		breadcrumbs.AddItem("/c/"+cat.Slug+"/1", cat.Name)
	}
	breadcrumbs.AddItem("", category.Name)

	products, err := s.productRepo.Search(ctx, category.ShortId(), dto.ItemsPerPage, (dto.Page-1)*dto.ItemsPerPage)
	if err != nil {
		return nil, err
	}

	return &CatalogCategoryModel{
		webPage:      webPage,
		categoryList: categoryList,
		breadcrumbs:  breadcrumbs,
		pagination:   pagination,
		listProducts: products,
	}, err
}
