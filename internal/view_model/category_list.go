package view_model

import "github.com/divilla/golastore/internal/domain_model"

type (
	CategoryList struct {
		CurrentCategorySlug string
		CurrentCategory     *domain_model.TaxonomyItem
		ListedCategory      *domain_model.TaxonomyItem
		URLBuilder          CategoryListURLBuilder
	}

	CategoryListURLBuilder func(slug string) string
)

func NewCategoryList(currentCategory *domain_model.TaxonomyItem, urlBuilder CategoryListURLBuilder) *CategoryList {
	c := &CategoryList{
		CurrentCategorySlug: currentCategory.Slug,
		CurrentCategory:     currentCategory,
		URLBuilder:          urlBuilder,
	}
	if len(currentCategory.Children) > 0 {
		c.ListedCategory = currentCategory
	} else {
		c.ListedCategory = currentCategory.Parent
	}

	return c
}
