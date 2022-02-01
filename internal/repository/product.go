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

func (r *Product) Search(ctx context.Context, search string, perPage, page int64) ([]*domain.ListProduct, int64, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Release()

	var total int64
	row := conn.QueryRow(ctx, `
		select product_fts_count($1) as total
	`, search)
	if row.Scan(&total) != nil {
		return nil, 0, err
	}

	pages := total / perPage
	if total%perPage > 0 {
		pages++
	}

	limit := perPage
	offset := (page - 1) * perPage
	rows, err := conn.Query(ctx, `
		select
			id,
			code,
			name,
			slug,
			old_price,
			price,
			description
		from
			product_fts_search($1, $2, $3);
	`, search, limit, offset)
	if err != nil {
		return nil, pages, err
	}
	defer rows.Close()

	var lps []*domain.ListProduct
	for rows.Next() {
		var lp domain.ListProduct
		if err = rows.Scan(&lp.Id, &lp.Code, &lp.Name, &lp.Slug, &lp.OldPrice, &lp.Price, &lp.Description); err != nil {
			return nil, pages, err
		}
		lps = append(lps, &lp)
	}

	return lps, pages, nil
}
