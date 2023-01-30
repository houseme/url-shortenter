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

var insRedisCache = iRedisCache{}

type iRedisCache struct {
	ctx context.Context
}

// RedisCache is the interface for redis cache
func RedisCache() *iRedisCache {
	return &insRedisCache
}

// ShortConn returns the redis connection
func (r *iRedisCache) ShortConn(_ context.Context) string {
	return shortConn
}

// ShortMirrorKey returns the redis mirror key
func (r *iRedisCache) ShortMirrorKey(_ context.Context, shortURL string) string {
	return shortMirrorKey + shortURL
}

// ShortMirrorQueue returns the redis mirror queue
func (r *iRedisCache) ShortMirrorQueue(_ context.Context) string {
	return shortMirrorQueue
}

// ShortAuditQueue returns the redis audit queue
func (r *iRedisCache) ShortAuditQueue(_ context.Context) string {
	return shortAuditQueue
}

// ShortRequestConn returns the redis request connection
func (r *iRedisCache) ShortRequestConn(_ context.Context) string {
	return shortRequestConn
}

// ShortCacheConn returns the redis cache connection
func (r *iRedisCache) ShortCacheConn(_ context.Context) string {
	return shortCacheConn
}

// ShortAccessLogQueue returns the redis access log queue
func (r *iRedisCache) ShortAccessLogQueue(_ context.Context) string {
	return shortAccessLogQueue
}

// ShortCacheObject returns the redis cache object
func (r *iRedisCache) ShortCacheObject(_ context.Context, shortURL string) string {
	return shortCacheObject + shortURL
}

// ShortAccessLogSummaryQueue returns the redis access log summary queue
func (r *iRedisCache) ShortAccessLogSummaryQueue(_ context.Context) string {
	return shortAccessLogSummaryQueue
}

// ShortAccessTokenConn returns the redis access token connection
func (r *iRedisCache) ShortAccessTokenConn(_ context.Context) string {
	return shortAccessTokenConn
}

// ShortAccessTokenKey returns the redis access token key
func (r *iRedisCache) ShortAccessTokenKey(_ context.Context, accessToken string) string {
	return shortAccessTokenKey + accessToken
}

// ShortAuthorizationKey returns the redis authorization key
func (r *iRedisCache) ShortAuthorizationKey(_ context.Context, authorization string) string {
	return shortAuthorizationKey + authorization
}
