/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package jwtmanager

import (
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// Cache in memory temporary store for responses from /validate for jwt
var Cache *cache.Cache
var expire int = 20 // default 20 minutes
var dExp time.Duration

func cacheConfigure() { _ = "STUB: not implemented"; return }

// log.Debugf("cacheConfigure expire %d dExp %d purgecheck %d", expire, dExp, purgeCheck)

// CachedResponse caches the JWT response
// type CachedResponse struct {
// 	*CaptureWriter
// 	rawResponse []byte
// }

// JWTCacheHandler looks for a JWT and...
// returns a cached response
// or passes the request to /validate
// all tests for JWTCacheHandler are present in `handlers/validate_test.go` to avoid circular imports
func JWTCacheHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// check to see if we have headers cached for this jwt

// found it in cache!

// TODO: instead of the copy for each, can we just append the whole blob?
// or better still can we just cache the entire response including 200OK?

// see responses.addErrandCancelRequest()
// r.Context().Done() is still open
// cache the response headers for this jwt
// log.Debug("setting cache for %+v", w.Header().Clone())

// log.Debugf("claims expire, time.now.unix, dExp %d - %d = %d > %d", claims.ExpiresAt, now, claims.ExpiresAt-now, int64(dExp))
// log.Debugf("time.Duration((claims.ExpiresAt-time.Now().Unix())*time.Second.Nanoseconds()) %d", time.Duration((claims.ExpiresAt-time.Now().Unix())*time.Second.Nanoseconds()))

// getCacheExpirationDuration - return time.Duration til the jwt should be purged from cache
// first see if the jwt's expiration will arrive before the cache expiration
// if this jwt expires in 10 minutes then we don't want to cache it for 20
// this might happen if the jwt expiration is set to 240 minutes, and the user last logged into the IdP 230 minutes ago
// then the user went away, cache was purged and now they return with 10 minutes left before token expiration
func getCacheExpirationDuration(claims *VouchClaims) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
