package repository

import (
	"context"
	"github.com/divilla/golastore/internal/domain"
	"github.com/divilla/golastore/pkg/postgres"
)

type (
	Product struct {
		pool *postgres.Pool
	}
)

func NewProductRepository(pool *postgres.Pool) *Product {
	return &Product{
		pool: pool,
	}
}

func (r *Taxonomy) Search(ctx context.Context, fts string, records, page int64) ([]*domain.TaxonomyItem, error) {
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
			path,
			position,
			parent_id,
			parent_slug
		from taxonomy_item_view
		where root or path is not null;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tis []*domain.TaxonomyItem
	for rows.Next() {
		var ti domain.TaxonomyItem
		if err = rows.Scan(&ti.Id, &ti.Name, &ti.Slug, &ti.Root, &ti.Path, &ti.Position, &ti.ParentId, &ti.ParentSlug); err != nil {
			return nil, err
		}
		tis = append(tis, &ti)
	}

	return tis, nil
}
