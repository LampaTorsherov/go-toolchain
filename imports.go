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

	// In-memory adapters
	_ "github.com/eko/gocache/store/bigcache/v4"
	_ "github.com/eko/gocache/store/freecache/v4"
	_ "github.com/eko/gocache/store/go_cache/v4"
	_ "github.com/eko/gocache/store/ristretto/v4"

	// Redis / Valkey adapters
	_ "github.com/eko/gocache/store/redis/v4"
	_ "github.com/eko/gocache/store/rediscluster/v4"
	_ "github.com/eko/gocache/store/rueidis/v4"
	_ "github.com/eko/gocache/store/valkey/v4"

	// Distributed adapters
	_ "github.com/eko/gocache/store/hazelcast/v4"
	_ "github.com/eko/gocache/store/memcache/v4"

	// Legacy adapters
	_ "github.com/eko/gocache/store/pegasus/v4"

	// In-memory clients
	_ "github.com/allegro/bigcache/v3"
	_ "github.com/coocood/freecache"
	_ "github.com/dgraph-io/ristretto/v2"
	_ "github.com/patrickmn/go-cache"

	// Redis / Valkey clients
	_ "github.com/redis/go-redis/v9"
	_ "github.com/redis/rueidis"
	_ "github.com/valkey-io/valkey-go"

	// Distributed clients
	_ "github.com/bradfitz/gomemcache/memcache"
	_ "github.com/hazelcast/hazelcast-go-client"

	// Legacy clients
	_ "github.com/XiaoMi/pegasus-go-client/pegasus"
)

func main() {}
