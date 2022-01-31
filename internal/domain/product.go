package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type (
	ListProduct struct {
		Id          uuid.UUID
		Name        string
		Slug        string
		OldPrice    pgtype.Numeric
		Price       pgtype.Numeric
		Description string
	}
)
