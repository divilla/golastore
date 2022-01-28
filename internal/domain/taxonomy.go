package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type (
	TaxonomyItem struct {
		Id         uuid.UUID
		Name       string
		Slug       string
		Root       bool
		Path       pgtype.Text
		SlugPath   pgtype.Text
		Position   pgtype.Int8
		ParentId   pgtype.UUID
		ParentSlug pgtype.Text
		Children   []*TaxonomyItem
	}
)
