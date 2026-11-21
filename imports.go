
package main

import (
	// Core
	_ "github.com/eko/gocache/lib/v4/cache"
	_ "github.com/eko/gocache/lib/v4/codec"
	_ "github.com/eko/gocache/lib/v4/marshaler"
	_ "github.com/eko/gocache/lib/v4/metrics"
	_ "github.com/eko/gocache/lib/v4/store"

	// Mocks
	_ "github.com/eko/gocache/lib/v4/mocks/cache"
	_ "github.com/eko/gocache/lib/v4/mocks/codec"
	_ "github.com/eko/gocache/lib/v4/mocks/metrics"
	_ "github.com/eko/gocache/lib/v4/mocks/store"

	// In-memory stores
	_ "github.com/eko/gocache/store/bigcache/v4"
	_ "github.com/eko/gocache/store/freecache/v4"
	_ "github.com/eko/gocache/store/go_cache/v4"
	_ "github.com/eko/gocache/store/ristretto/v4"

	// Redis / Valkey stores
	_ "github.com/eko/gocache/store/redis/v4"
	_ "github.com/eko/gocache/store/rediscluster/v4"
	_ "github.com/eko/gocache/store/rueidis/v4"
	_ "github.com/eko/gocache/store/valkey/v4"

	// Distributed stores
	_ "github.com/eko/gocache/store/hazelcast/v4"
	_ "github.com/eko/gocache/store/memcache/v4"

	// Legacy stores
	_ "github.com/eko/gocache/store/pegasus/v4"
)

func main() {}
