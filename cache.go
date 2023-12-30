package cache

import (
	"context"
	"sync"
)

type data[V any] struct {
	value V
	error error
	done  <-chan struct{}
}

type Cache[K comparable, V any] struct {
	mu    sync.Mutex
	cache map[K]*data[V]
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{cache: map[K]*data[V]{}}
}

func (c *Cache[K, V]) Compute(ctx context.Context, key K, fn func(context.Context, K) (V, error)) (V, error) {
	c.mu.Lock()
	if d, ok := c.cache[key]; ok {
		c.mu.Unlock()
		<-d.done
		return d.value, d.error
	}

	done := make(chan struct{})
	d := &data[V]{done: done}
	c.cache[key] = d
	c.mu.Unlock()

	v, err := fn(ctx, key)
	d.value = v
	d.error = err
	close(done)
	return v, err
}
