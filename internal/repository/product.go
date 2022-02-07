package repository

import (
	"context"
	"github.com/divilla/golastore/internal/domain_model"
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

func (r *Product) SearchCount(ctx context.Context, search string) (int64, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var total int64
	row := conn.QueryRow(ctx, `
		select
			cnt
		from
			product_fts_count($1);
	`, search)
	err = row.Scan(&total)

	return total, err
}

func (r *Product) Search(ctx context.Context, search string, limit, offset int64) ([]*domain_model.ListProduct, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

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
		return nil, err
	}
	defer rows.Close()

	var lps []*domain_model.ListProduct
	for rows.Next() {
		var lp domain_model.ListProduct
		if err = rows.Scan(&lp.Id, &lp.Code, &lp.Name, &lp.Slug, &lp.OldPrice, &lp.Price, &lp.Description); err != nil {
			return nil, err
		}
		lps = append(lps, &lp)
	}

	return lps, nil
}
