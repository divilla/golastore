package repository

import (
	"context"
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/divilla/golastore/pkg/postgres"
)

type (
	Taxonomy struct {
		pool *postgres.Pool
	}
)

func NewTaxonomyRepository(pool *postgres.Pool) *Taxonomy {
	return &Taxonomy{
		pool: pool,
	}
}

func (r *Taxonomy) All(ctx context.Context) ([]*domain_model.TaxonomyItem, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `
		select
			id,
			name,
			slug,
			root,
			properties,
			parent_id,
			parent_slug,
			position
		from taxonomy_item_details;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tis []*domain_model.TaxonomyItem
	for rows.Next() {
		var ti domain_model.TaxonomyItem
		if err = rows.Scan(&ti.Id, &ti.Name, &ti.Slug, &ti.Root, &ti.Properties, &ti.ParentId, &ti.ParentSlug, &ti.Position); err != nil {
			return nil, err
		}
		tis = append(tis, &ti)
	}

	return tis, nil
}
