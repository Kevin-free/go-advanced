package dao

import (
	"context"

	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	cache *fanout.Fanout
}

// New new a dao and return.
func New() (d Dao, err error) {
	d = &dao{
		cache: fanout.New("cache"),
	}
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
