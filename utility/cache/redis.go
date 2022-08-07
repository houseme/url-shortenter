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

	shortAccessTokenConn = "short_access_token"

	shortAccessTokenKey = "short_access_token_key_"

	shortAuthorizationKey = "short_authorization_key_"
)

var insRedisCache = iRedisCache{
	ctx: context.Background(),
}

type iRedisCache struct {
	ctx context.Context
}

// RedisCache is the interface for redis cache
func RedisCache() *iRedisCache {
	return &insRedisCache
}

// ShortConn returns the redis connection
func (r *iRedisCache) ShortConn(ctx context.Context) string {
	r.ctx = ctx
	return shortConn
}

// ShortMirrorKey returns the redis mirror key
func (r *iRedisCache) ShortMirrorKey(ctx context.Context, shortURL string) string {
	r.ctx = ctx
	return shortMirrorKey + shortURL
}

// ShortMirrorQueue returns the redis mirror queue
func (r *iRedisCache) ShortMirrorQueue(ctx context.Context) string {
	r.ctx = ctx
	return shortMirrorQueue
}

// ShortAuditQueue returns the redis audit queue
func (r *iRedisCache) ShortAuditQueue(ctx context.Context) string {
	r.ctx = ctx
	return shortAuditQueue
}

// ShortRequestConn returns the redis request connection
func (r *iRedisCache) ShortRequestConn(ctx context.Context) string {
	r.ctx = ctx
	return shortRequestConn
}

// ShortCacheConn returns the redis cache connection
func (r *iRedisCache) ShortCacheConn(ctx context.Context) string {
	r.ctx = ctx
	return shortCacheConn
}

// ShortAccessLogQueue returns the redis access log queue
func (r *iRedisCache) ShortAccessLogQueue(ctx context.Context) string {
	r.ctx = ctx
	return shortAccessLogQueue
}

// ShortCacheObject returns the redis cache object
func (r *iRedisCache) ShortCacheObject(ctx context.Context, shortURL string) string {
	r.ctx = ctx
	return shortCacheObject + shortURL
}

// ShortAccessLogSummaryQueue returns the redis access log summary queue
func (r *iRedisCache) ShortAccessLogSummaryQueue(ctx context.Context) string {
	r.ctx = ctx
	return shortAccessLogSummaryQueue
}

// ShortAccessTokenConn returns the redis access token connection
func (r *iRedisCache) ShortAccessTokenConn(ctx context.Context) string {
	r.ctx = ctx
	return shortAccessTokenConn
}

// ShortAccessTokenKey returns the redis access token key
func (r *iRedisCache) ShortAccessTokenKey(ctx context.Context, accessToken string) string {
	r.ctx = ctx
	return shortAccessTokenKey + accessToken
}

// ShortAuthorizationKey returns the redis authorization key
func (r *iRedisCache) ShortAuthorizationKey(ctx context.Context, authorization string) string {
	r.ctx = ctx
	return shortAuthorizationKey + authorization
}
