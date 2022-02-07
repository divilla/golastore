package view_model

import "github.com/divilla/golastore/internal/domain_model"

type (
	CategoryVM struct {
		CurrentCategorySlug string
		CurrentCategory     *domain_model.TaxonomyItem
		ListedCategory      *domain_model.TaxonomyItem
	}
)

func NewCategoryVM(currentCategory *domain_model.TaxonomyItem) *CategoryVM {
	c := &CategoryVM{
		CurrentCategorySlug: currentCategory.Slug,
		CurrentCategory:     currentCategory,
	}
	if len(currentCategory.Children) > 0 {
		c.ListedCategory = currentCategory
	} else {
		c.ListedCategory = currentCategory.Parent
	}

	return c
}
