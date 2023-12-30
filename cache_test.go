package cache

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	c := New[string, int]()

	count := uint64(0)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out, err := c.Compute(context.Background(), "a", func(ctx context.Context, key string) (int, error) {
				atomic.AddUint64(&count, 1)
				time.Sleep(100 * time.Millisecond)
				return 99, nil
			})
			if err != nil {
				t.Error()
			}
			if out != 99 {
				t.Error()
			}
		}()
	}
	wg.Wait()
	if count != 1 {
		t.Fatal()
	}
}
