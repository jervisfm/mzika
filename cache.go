package mzika

// Author: Jervis Muindi
// Date: December 2015

import (
	"strings"
)

// Caching entity
type CachedVideoJson struct {
	// A video ID string, e.g. UScjy1431460
	VideoId string
	// The cached Response from VEVO API
	Response VideoJSON
}

var cacheStore map[string]CachedVideoJson

// Initialized the cache store to a non-nil value. If already initialized
// then this is a no-op and has no effect.
func initializeCacheStore() {
	if cacheStore == nil {
		cacheStore = make(map[string]CachedVideoJson)
	}
}

// Takes given VideoJSON |input| and saves it into our in-memory datastore
// using the key |cachekey|. This is so that the response can quickly
// be looked up in the future. |r| is the Http request for the current session.
func CacheVideoJsonResponse(input VideoJSON, cacheKey string) (err error) {
	cacheKey = strings.ToLower(cacheKey)
	cachedData := CachedVideoJson{
		VideoId:  cacheKey,
		Response: input,
	}

	// Try to determine if cache entry already exists
	cachedData, ok := cacheStore[cacheKey]
	if !ok {
		println("Cache key no exist")
		// Cache entry does not exist, so we should store it.
		initializeCacheStore()
		cacheStore[cacheKey] = cachedData
	}
	return nil
}

// Attempts to find a cached response given |videoid|. Note that |videoid| should
// be the primary key we use to store the cached responses. Returns true iff a valid
// response was found.
func GetCachedVideoResponse(videoid string) (found bool, response *VideoJSON, err error) {
	cacheKey := videoid
	// Note: In Go, it's OK to read from a nil map. It's the task of
	// writing to it that's problematic and causes a panic.
	cachedData, ok := cacheStore[cacheKey]
	if ok {
		// Cached response found
		// TODO(jervis): Figure out why cache response returned is currenyl empty.
		return true, &cachedData.Response, nil
	} else {
		// No cached video response found.
		return false, nil, err
	}

}
