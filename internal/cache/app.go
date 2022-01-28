package cache

import (
	"sync"
)

type (
	App struct {
		data map[string]interface{}
		rwm  *sync.RWMutex
	}
)

func NewAppCache() *App {
	data := map[string]interface{}{
		"title": "Golastore",
	}

	return &App{
		data: data,
		rwm:  new(sync.RWMutex),
	}
}

func (c *App) Title() string {
	return c.data["title"].(string)
}
