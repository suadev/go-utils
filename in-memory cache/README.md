Basic TTL supported in-memory cache implementation without using any other 3rd party package.

Usage:

```go
	ttlCache, err := NewTTLCache(time.Minutes * 30)

    err = ttlCache.Set("cachekey", <value>)

    value := ttlCache.Get("cacheKey")
```