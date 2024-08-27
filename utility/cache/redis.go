// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

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

// IRedisCache is the interface for Redis cache
type IRedisCache struct {
	ctx context.Context
}

// RedisCache is the interface for Redis cache
func RedisCache() *IRedisCache {
	return &IRedisCache{}
}

// ShortConn returns the Redis connection
func (r *IRedisCache) ShortConn(_ context.Context) string {
	return shortConn
}

// ShortMirrorKey returns the Redis mirror key
func (r *IRedisCache) ShortMirrorKey(_ context.Context, shortURL string) string {
	return shortMirrorKey + shortURL
}

// ShortMirrorQueue returns the Redis mirror queue
func (r *IRedisCache) ShortMirrorQueue(_ context.Context) string {
	return shortMirrorQueue
}

// ShortAuditQueue returns the Redis audit queue
func (r *IRedisCache) ShortAuditQueue(_ context.Context) string {
	return shortAuditQueue
}

// ShortRequestConn returns the Redis request connection
func (r *IRedisCache) ShortRequestConn(_ context.Context) string {
	return shortRequestConn
}

// ShortCacheConn returns the Redis cache connection
func (r *IRedisCache) ShortCacheConn(_ context.Context) string {
	return shortCacheConn
}

// ShortAccessLogQueue returns the Redis access log queue
func (r *IRedisCache) ShortAccessLogQueue(_ context.Context) string {
	return shortAccessLogQueue
}

// ShortCacheObject returns the Redis cache object
func (r *IRedisCache) ShortCacheObject(_ context.Context, shortURL string) string {
	return shortCacheObject + shortURL
}

// ShortAccessLogSummaryQueue returns the Redis access log summary queue
func (r *IRedisCache) ShortAccessLogSummaryQueue(_ context.Context) string {
	return shortAccessLogSummaryQueue
}

// ShortAccessTokenConn returns the Redis access token connection
func (r *IRedisCache) ShortAccessTokenConn(_ context.Context) string {
	return shortAccessTokenConn
}

// ShortAccessTokenKey returns the Redis access token key
func (r *IRedisCache) ShortAccessTokenKey(_ context.Context, accessToken string) string {
	return shortAccessTokenKey + accessToken
}

// ShortAuthorizationKey returns the Redis authorization key
func (r *IRedisCache) ShortAuthorizationKey(_ context.Context, authorization string) string {
	return shortAuthorizationKey + authorization
}
