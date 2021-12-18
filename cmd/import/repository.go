package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"regexp"
	"strings"
)

type (
	repository struct {
		ctx context.Context
		pool *pgxpool.Pool
	}

	properties struct {
		breadcrumbPath []breadcrumb
	}

	breadcrumb struct {
		slug string
		name string
	}
)

func newRepository(ctx context.Context, pool *pgxpool.Pool) *repository {
	return &repository{
		ctx:  ctx,
		pool: pool,
	}
}

func (r *repository) categoryNameToSlug() {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	connExec, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer connExec.Release()

	rows, err := conn.Query(r.ctx, "select id, name from taxonomy_item")
	fatal(err)
	defer rows.Close()

	regWord := regexp.MustCompile(`[^\w]`)
	regSpace := regexp.MustCompile(`\s+`)
	for rows.Next() {
		var id uuid.UUID
		var name string
		err = rows.Scan(&id, &name)
		fatal(err)

		byteId, err := id.MarshalBinary()
		fatal(err)

		urlId := base64.RawURLEncoding.EncodeToString(byteId)

		slug := strings.ToLower(name)
		slug = strings.ReplaceAll(slug, "č", "c")
		slug = strings.ReplaceAll(slug, "ć", "c")
		slug = strings.ReplaceAll(slug, "đ", "d")
		slug = strings.ReplaceAll(slug, "š", "s")
		slug = strings.ReplaceAll(slug, "ž", "z")

		slug = regWord.ReplaceAllString(slug, " ")
		slug = regSpace.ReplaceAllString(slug, " ")
		slug = strcase.ToKebab(slug)

		slug = slug + "/" + urlId

		_, err = connExec.Exec(r.ctx, "update taxonomy_item set slug=$1 where id=$2", slug, id)
		fatal(err)

		fmt.Println(urlId, slug)
	}
}

func (r *repository) categoryBreadcrumbPath() {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	connExec, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer connExec.Release()

	rows, err := conn.Query(r.ctx, "select id, slug, parent_id from taxonomy_item")
	fatal(err)
	defer rows.Close()

	for rows.Next() {
		var id, parentId uuid.UUID
		var slug string
		err = rows.Scan(&id, &slug, &parentId)
		fatal(err)

		//byteId, err := id.MarshalBinary()
		//fatal(err)
		//
		//urlId := base64.RawURLEncoding.EncodeToString(byteId)
		//
		//slug := strings.ToLower(name)
		//slug = strings.ReplaceAll(slug, "č", "c")
		//slug = strings.ReplaceAll(slug, "ć", "c")
		//slug = strings.ReplaceAll(slug, "đ", "d")
		//slug = strings.ReplaceAll(slug, "š", "s")
		//slug = strings.ReplaceAll(slug, "ž", "z")
		//
		//slug = regWord.ReplaceAllString(slug, " ")
		//slug = regSpace.ReplaceAllString(slug, " ")
		//slug = strcase.ToKebab(slug)
		//
		//slug = slug + "-" + urlId
		//
		//_, err = connExec.Exec(r.ctx, "update taxonomy_item set slug=$1 where id=$2", slug, id)
		//fatal(err)
		//
		//fmt.Println(urlId, slug)
	}
}

func (r *repository) resetImportProductActive() {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(r.ctx, `update import_product set active=false`)
	fatal(err)
}

func (r *repository) truncateProductBreadcrumbs() {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(r.ctx, `truncate table import_product_breadcrumb`)
	fatal(err)
}

func (r *repository) listInactiveUrls() []string {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "select url from import_product where active=false")
	fatal(err)
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		fatal(err)

		urls = append(urls, url)
	}

	return urls
}

func (r *repository) truncateProductProductProperties() {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(r.ctx, `truncate table import_product_product_property`)
	fatal(err)
}

func (r *repository) updateProduct(name string, oldPrice, price pgtype.Numeric, description, productCode string) {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "update import_product set name=$1, old_price=$2, price=$3, description=$4 where code=$5",
		name, oldPrice, price, description, productCode)
	fatal(err)
}

func (r *repository) updateProductImageCount(count int, productCode string) {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "update import_product set image_count=$1 where code=$2",
		count, productCode)
	fatal(err)
}

func (r *repository) insertProductBreadcrumb(productCode, categoryCode string, pos int, url, name string) {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "insert into import_product_breadcrumb (product_code, category_code, pos, url, name) values ($1, $2, $3, $4, $5)",
		productCode, categoryCode, pos, url, name)
	warn(err)
}

func (r *repository) insertProductProductProperty(productCode, propCode, propValue string) {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, _ = conn.Exec(r.ctx, "insert into import_product_property (code) values ($1)", propCode)

	_, err = conn.Exec(r.ctx,
		"insert into import_product_product_property (product_code, product_property_code, product_property_value) values ($1, $2, $3)",
		productCode, propCode, propValue)
	warn(err)
}

func (r repository) setProductActive(productCode string) {
	conn, err := r.pool.Acquire(r.ctx)
	fatal(err)
	defer conn.Release()

	_, err = conn.Exec(r.ctx, "update import_product set active=true where code=$1", productCode)
	fatal(err)
}
