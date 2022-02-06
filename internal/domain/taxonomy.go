package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	TaxonomyItem struct {
		Id         uuid.UUID
		Name       string
		Slug       string
		Root       bool
		Properties TaxonomyProperties
		Path       []*TaxonomyPathItem
		Position   pgtype.Int8
		ParentId   pgtype.UUID
		ParentSlug pgtype.Text
		Children   []*TaxonomyItem
	}

	TaxonomyPathItem struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	}

	TaxonomyProperties struct {
		TotalProducts int64 `json:"totalProducts"`
	}
)

func (t *TaxonomyItem) TotalProducts() string {
	p := message.NewPrinter(language.BritishEnglish)
	return p.Sprintf("%d", t.Properties.TotalProducts)
}
