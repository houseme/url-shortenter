package cache

import (
	"context"
)

const (
	shortConn = "short_url"

	shortMirrorQueue = "short_mirror_queue"

	shortMirrorKey = "short_mirror_key_"

	shortAuditQueue = "short_audit_queue"

	shortRequestConn = "short"

	shortCacheConn = "short_cache"

	shortAccessLogQueue = "short_access_log_queue"

	shortCacheObject = "short_cache_object_"

	shortAccessLogSummaryQueue = "short_access_log_summary_queue"
)

var insRedisCache = iRedisCache{}

type iRedisCache struct {
}

// RedisCache is the interface for redis cache
func RedisCache() *iRedisCache {
	return &insRedisCache
}

// ShortConn returns the redis connection
func (r *iRedisCache) ShortConn(ctx context.Context) string {
	return shortConn
}

// ShortMirrorKey returns the redis mirror key
func (r *iRedisCache) ShortMirrorKey(ctx context.Context, shortURL string) string {
	return shortMirrorKey + shortURL
}

// ShortMirrorQueue returns the redis mirror queue
func (r *iRedisCache) ShortMirrorQueue(ctx context.Context) string {
	return shortMirrorQueue
}

// ShortAuditQueue returns the redis audit queue
func (r *iRedisCache) ShortAuditQueue(ctx context.Context) string {
	return shortAuditQueue
}

// ShortRequestConn returns the redis request connection
func (r *iRedisCache) ShortRequestConn(ctx context.Context) string {
	return shortRequestConn
}

// ShortCacheConn returns the redis cache connection
func (r *iRedisCache) ShortCacheConn(ctx context.Context) string {
	return shortCacheConn
}

// ShortAccessLogQueue returns the redis access log queue
func (r *iRedisCache) ShortAccessLogQueue(ctx context.Context) string {
	return shortAccessLogQueue
}

// ShortCacheObject returns the redis cache object
func (r *iRedisCache) ShortCacheObject(ctx context.Context, shortURL string) string {
	return shortCacheObject + shortURL
}

// ShortAccessLogSummaryQueue returns the redis access log summary queue
func (r *iRedisCache) ShortAccessLogSummaryQueue(ctx context.Context) string {
	return shortAccessLogSummaryQueue
}
