# cache

This is a golang in-meomry cache library
which prevents [cache stampedes](https://en.wikipedia.org/wiki/Cache_stampede).

# Usage

```go
import "github.com/skaji/go-cache"

c := cache.New[string, int]()

// we can make sure database.Query with the same key must be called at most once
count, err := c.Compute(context.Background(), "key", func(ctx context.Context, key string) (int, error) {
    rows, err := database.Query(ctx, key)
    if err != nil {
        return 0, err
    }
    return len(rows), nil
})
```
