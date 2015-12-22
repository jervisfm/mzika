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
		// Cache entry does not exist, so we should store it.
		cacheStore[cacheKey] = cachedData
	}
	return nil
}

// Attempts to find a cached response given |videoid|. Note that |videoid| should
// be the primary key we use to store the cached responses. Returns true iff a valid
// response was found.
func GetCachedVideoResponse(videoid string) (found bool, response *VideoJSON, err error) {
	cacheKey := videoid
	cachedData, ok := cacheStore[cacheKey]
	if ok {
		// Cached response found
		return true, &cachedData.Response, nil
	} else {
		// No cached video response found.
		return false, nil, err
	}

}
