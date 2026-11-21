
module offline-cache

go 1.27

require (
	github.com/eko/gocache/lib/v4 v4.4.0

	// In-memory
	github.com/eko/gocache/store/bigcache/v4 v4.2.10
	github.com/eko/gocache/store/freecache/v4 v4.2.10
	github.com/eko/gocache/store/go_cache/v4 v4.2.10
	github.com/eko/gocache/store/ristretto/v4 v4.3.8

	// Redis / Valkey
	github.com/eko/gocache/store/redis/v4 v4.2.13
	github.com/eko/gocache/store/rediscluster/v4 v4.2.10
	github.com/eko/gocache/store/rueidis/v4 v4.1.14
	github.com/eko/gocache/store/valkey/v4 v4.2.11

	// Distributed cache
	github.com/eko/gocache/store/hazelcast/v4 v4.2.10
	github.com/eko/gocache/store/memcache/v4 v4.2.11
	
	// Legacy
	github.com/eko/gocache/store/pegasus/v4 v4.2.6
)
