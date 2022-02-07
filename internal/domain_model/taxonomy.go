package domain_model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

type (
	TaxonomyItem struct {
		Id         uuid.UUID
		Name       string
		Slug       string
		Root       bool
		Properties TaxonomyProperties
		ParentId   pgtype.UUID
		ParentSlug pgtype.Text
		Position   pgtype.Int8
		Parent     *TaxonomyItem
		Path       []*TaxonomyItem
		Children   []*TaxonomyItem
	}

	TaxonomyProperties struct {
		TotalProducts int64 `json:"totalProducts"`
	}
)

func (t *TaxonomyItem) ShortId() string {
	return strings.ReplaceAll(t.Id.String(), "-", "")
}

func (t *TaxonomyItem) TotalProducts() string {
	p := message.NewPrinter(language.German)
	return p.Sprintf("%d", t.Properties.TotalProducts)
}
