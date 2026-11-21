module offline-cache

go 1.27

require (
	// Core
	github.com/eko/gocache/lib/v4 v4.4.0

	// In-memory adapters
	github.com/eko/gocache/store/bigcache/v4 v4.2.10
	github.com/eko/gocache/store/freecache/v4 v4.2.10
	github.com/eko/gocache/store/go_cache/v4 v4.2.10
	github.com/eko/gocache/store/ristretto/v4 v4.3.8

	// Redis / Valkey adapters
	github.com/eko/gocache/store/redis/v4 v4.2.13
	github.com/eko/gocache/store/rediscluster/v4 v4.2.10
	github.com/eko/gocache/store/rueidis/v4 v4.1.14
	github.com/eko/gocache/store/valkey/v4 v4.2.11

	// Distributed adapters
	github.com/eko/gocache/store/hazelcast/v4 v4.2.10
	github.com/eko/gocache/store/memcache/v4 v4.2.11

	// Legacy adapters
	github.com/eko/gocache/store/pegasus/v4 v4.2.6

	// In-memory clients
	github.com/allegro/bigcache/v3 v3.2.0
	github.com/coocood/freecache v1.2.7
	github.com/dgraph-io/ristretto/v2 v2.4.2
	github.com/patrickmn/go-cache v2.1.0+incompatible

	// Redis / Valkey clients
	github.com/redis/go-redis/v9 v9.23.0
	github.com/redis/rueidis v1.0.80
	github.com/valkey-io/valkey-go v1.0.80

	// Distributed clients
	github.com/bradfitz/gomemcache v0.0.0-20260422231931-4d751bb6e37c
	github.com/hazelcast/hazelcast-go-client v1.6.0

	// Legacy clients
	github.com/XiaoMi/pegasus-go-client v0.0.0-20220519103347-ba0e68465cd5
)
