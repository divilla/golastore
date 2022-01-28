package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	pool *pgxpool.Pool
}

func NewPool(dsn string) *Pool {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 100

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return &Pool{
		pool: pool,
	}
}

func (p *Pool) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return p.pool.Acquire(ctx)
}

func (p *Pool) Close() {
	p.Close()
}
