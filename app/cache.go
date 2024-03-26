package app

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"grest.dev/grest"
)

// Cache returns a pointer to the cacheUtil instance (cache).
// If cache is not initialized, it creates a new cacheUtil instance, configures it, and assigns it to cache.
// It ensures that only one instance of cacheUtil is created and reused.
func Cache() *cacheUtil {
	if cache == nil {
		cache = &cacheUtil{}
		cache.configure()
	}
	return cache
}

// cache is a pointer to a cacheUtil instance.
// It is used to store and access the singleton instance of cacheUtil.
var cache *cacheUtil

// cacheUtil represents a cache utility.
// It embeds grest.Cache, indicating that cacheUtil inherits from grest.Cache.
type cacheUtil struct {
	grest.Cache
}

// configure configures the cache utility instance.
// It sets the expiration time (c.Exp) to 24 hours and initializes the Redis client (c.RedisClient) with the provided Redis options.
// It sets the context (c.Ctx) to the background context.
// It pings the Redis server to check the connection status and stores the result in the err variable.
// If there is an error connecting to Redis, it logs the error and the Redis connection details.
// Otherwise, it sets c.IsUseRedis to true and logs a successful cache configuration with Redis.
func (c *cacheUtil) configure() {
	c.Exp = 24 * time.Hour
	c.RedisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Username: REDIS_USERNAME,
		Password: REDIS_PASSWORD,
		DB:       REDIS_CACHE_DB,
	})
	c.Ctx = context.Background()
	err := c.RedisClient.Ping(c.Ctx).Err()
	if err != nil {
		Logger().Error().
			Err(err).
			Str("REDIS_HOST", REDIS_HOST).
			Str("REDIS_PORT", REDIS_PORT).
			Str("REDIS_USERNAME", REDIS_USERNAME).
			Str("REDIS_PASSWORD", REDIS_PASSWORD).
			Int("REDIS_CACHE_DB", REDIS_CACHE_DB).
			Msg("Failed to connect to redis. The cache will be use in-memory local storage.")
	} else {
		c.IsUseRedis = true
		Logger().Info().Msg("Cache configured with redis.")
	}
}
