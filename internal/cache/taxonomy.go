package cache

import (
	"context"
	"errors"
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/divilla/golastore/internal/repository"
	"github.com/jackc/pgtype"
	"sync"
)

var TaxNotFoundErr = errors.New("taxonomy item not found")

type (
	Taxonomy struct {
		cache      map[string]*domain_model.TaxonomyItem
		repository *repository.Taxonomy
		rwm        *sync.RWMutex
	}
)

func NewTaxonomyCache(r *repository.Taxonomy) *Taxonomy {
	t := &Taxonomy{
		cache:      make(map[string]*domain_model.TaxonomyItem),
		repository: r,
		rwm:        new(sync.RWMutex),
	}
	if err := t.Refresh(context.Background()); err != nil {
		panic(err)
	}

	return t
}

func (c *Taxonomy) Get(key string) (*domain_model.TaxonomyItem, error) {
	c.rwm.RLock()
	defer c.rwm.RUnlock()

	if val, ok := c.cache[key]; ok {
		return val, nil
	} else {
		return nil, TaxNotFoundErr
	}
}

func (c *Taxonomy) Refresh(ctx context.Context) error {
	c.rwm.Lock()
	defer c.rwm.Unlock()

	tis, err := c.repository.All(ctx)
	if err != nil {
		return err
	}

	for key := range c.cache {
		delete(c.cache, key)
	}

	for _, ti := range tis {
		c.cache[ti.Slug] = ti
	}

	for _, ti := range tis {
		if ti.ParentSlug.Status == pgtype.Present {
			parent := c.cache[ti.ParentSlug.String]
			ti.Parent = parent
			ti.Path = c.path(ti.Parent)
			parent.Children = append(parent.Children, ti)
		}
	}

	return nil
}

func (c *Taxonomy) path(ti *domain_model.TaxonomyItem) []*domain_model.TaxonomyItem {
	if ti.ParentSlug.Status == pgtype.Present {
		parent := c.cache[ti.ParentSlug.String]
		return append(c.path(parent), ti)
	}

	return []*domain_model.TaxonomyItem{ti}
}
