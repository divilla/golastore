package maintenance_service

import (
	"context"
	"github.com/divilla/golastore/pkg/postgres"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/tidwall/sjson"
	"regexp"
	"strconv"
	"strings"
)

type (
	MaintenanceService struct {
		pool *postgres.Pool
	}
)

func New(pool *postgres.Pool) *MaintenanceService {
	return &MaintenanceService{
		pool: pool,
	}
}

func (s *MaintenanceService) RebuildTaxonomySlugs(ctx context.Context) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "update taxonomy_item set slug=null")
	if err != nil {
		return err
	}

	return s.rebuildTaxonomySlugsRecursive(ctx, uuid.MustParse("e9e73faa-7c57-11ec-a98d-0242ac110002"))
}

func (s *MaintenanceService) rebuildTaxonomySlugsRecursive(ctx context.Context, parentId uuid.UUID) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	connExec, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer connExec.Release()

	rows, err := conn.Query(ctx, `
		select id, name
		from taxonomy_item ti
			inner join taxonomy_item_parent tip on ti.id = tip.child_id
		where tip.parent_id = $1;
	`, parentId)
	if err != nil {
		return err
	}
	defer rows.Close()

	regWord := regexp.MustCompile(`[^\w]`)
	regSpace := regexp.MustCompile(`\s+`)
	for rows.Next() {
		var id uuid.UUID
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return err
		}

		slug := strings.ToLower(name)
		slug = strings.ReplaceAll(slug, "č", "c")
		slug = strings.ReplaceAll(slug, "ć", "c")
		slug = strings.ReplaceAll(slug, "đ", "d")
		slug = strings.ReplaceAll(slug, "š", "s")
		slug = strings.ReplaceAll(slug, "ž", "z")

		slug = regWord.ReplaceAllString(slug, " ")
		slug = regSpace.ReplaceAllString(slug, " ")
		slug = strcase.ToKebab(slug)

		_, err = connExec.Exec(ctx, "update taxonomy_item set slug=$1 where id=$2", slug, id)
		if err != nil {
			for i := 1; i < 30; i++ {
				_, err = connExec.Exec(ctx, "update taxonomy_item set slug=$1 where id=$2", slug+"-"+strconv.Itoa(i), id)
				if err == nil {
					break
				}
			}
			if err != nil {
				return err
			}
		}

		err = s.rebuildTaxonomySlugsRecursive(ctx, id)
	}

	return nil
}

func (s *MaintenanceService) RebuildTaxonomyParents(ctx context.Context) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	connExec, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer connExec.Release()

	_, err = conn.Exec(ctx, `
		update taxonomy_item_parent set path=null;
	`)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `
		update taxonomy_item_parent
		set path=jsonb_build_array(jsonb_build_object('slug', til.slug, 'name', til.name), jsonb_build_object('slug', tir.slug, 'name', tir.name))
		from taxonomy_item til
			inner join taxonomy_item_parent tip on til.id = tip.parent_id
			inner join taxonomy_item tir on tip.child_id = tir.id
		where taxonomy_item_parent.child_id = tir.id and til.root;
	`)
	if err != nil {
		return err
	}

	rows, err := conn.Query(ctx, `
		select
		    tip.parent_id,
			tip.child_id
		from taxonomy_item til
			inner join taxonomy_item_parent tip on til.id = tip.parent_id
			inner join taxonomy_item tir on tip.child_id = tir.id
		where til.root
		order by tir.name;
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var i int64
	for rows.Next() {
		var pid, cid uuid.UUID
		err = rows.Scan(&pid, &cid)
		if err != nil {
			return err
		}

		i++
		_, err = connExec.Exec(context.Background(), "update taxonomy_item_parent set position=$1 where parent_id=$2 and child_id=$3", i, pid, cid)

		if err = s.rebuildTaxonomyParentsRecursive(ctx, pid, cid); err != nil {
			return err
		}
	}

	return nil
}

func (s *MaintenanceService) rebuildTaxonomyParentsRecursive(ctx context.Context, parentId uuid.UUID, childId uuid.UUID) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	connExec, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer connExec.Release()

	rows, err := conn.Query(ctx, `
		select
			tipc.parent_id,
			tipc.child_id,
			tipp.path::text as path,
			ti.name,
			ti.slug
		from taxonomy_item_parent tipp
			inner join taxonomy_item_parent tipc on tipp.child_id = tipc.parent_id
			inner join taxonomy_item ti on tipc.child_id = ti.id
		where tipp.parent_id=$1 and tipp.child_id=$2
		order by ti.name;
	`, parentId, childId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var i int64
	for rows.Next() {
		var pid, cid uuid.UUID
		var path, name, slug string
		err = rows.Scan(&pid, &cid, &path, &name, &slug)
		if err != nil {
			return err
		}

		item, _ := sjson.Set("{}", "name", name)
		item, _ = sjson.Set(item, "slug", slug)
		path, _ = sjson.SetRaw(path, "-1", item)

		i++
		_, err = connExec.Exec(ctx, "update taxonomy_item_parent set path=$1::jsonb, position=$2 where parent_id=$3 and child_id=$4", path, i, pid, cid)
		if err != nil {
			return err
		}

		err = s.rebuildTaxonomyParentsRecursive(ctx, pid, cid)
	}

	return nil
}
