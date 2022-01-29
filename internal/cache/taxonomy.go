package cache

import (
	"context"
	"github.com/divilla/golastore/internal/domain"
	"github.com/divilla/golastore/internal/repository"
	"sync"
)

type (
	Taxonomy struct {
		cache      map[string]*domain.TaxonomyItem
		repository *repository.Taxonomy
		rwm        *sync.RWMutex
	}
)

func NewTaxonomyCache(r *repository.Taxonomy) *Taxonomy {
	t := &Taxonomy{
		cache:      make(map[string]*domain.TaxonomyItem),
		repository: r,
		rwm:        new(sync.RWMutex),
	}
	if err := t.Refresh(context.Background()); err != nil {
		panic(err)
	}

	return t
}

func (c *Taxonomy) ProductCategories() *domain.TaxonomyItem {
	c.rwm.RLock()
	defer c.rwm.RUnlock()

	return c.cache["product-categories-root"]
}

func (c *Taxonomy) Get(key string) *domain.TaxonomyItem {
	c.rwm.RLock()
	defer c.rwm.RUnlock()

	return c.cache[key]
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
		var parents []*domain.TaxonomyItem
		parentSlug := ti.ParentSlug.String
		for {
			if parentSlug == "" {
				break
			}
			parents = append(parents, c.cache[parentSlug])
			parentSlug = c.cache[parentSlug].ParentSlug.String
		}

		if l := len(parents); l > 0 {
			ti.Parents = make([]*domain.TaxonomyItem, l)
			l--
			for key, val := range parents {
				ti.Parents[l-key] = val
			}
		}

		if parent, ok := c.cache[ti.ParentSlug.String]; ok {
			parent.Children = append(parent.Children, ti)
		}
	}

	return nil
}
