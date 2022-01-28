package service

import (
	"github.com/divilla/golastore/internal/domain"
	"github.com/divilla/golastore/internal/repository"
)

type (
	Taxonomy struct {
		repository *repository.Taxonomy
		cache      map[string]domain.TaxonomyItem
	}
)

func NewTaxonomyService(r *repository.Taxonomy) *Taxonomy {
	return &Taxonomy{
		repository: r,
	}
}
