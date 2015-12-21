package mzika

// Author: Jervis Muindi
// Date: December 2015

import (
	"fmt"
	"net/http"
	"strings"
)

// Caching entity
type CachedVideoJson struct {
	// A video ID string, e.g. UScjy1431460
	VideoId string
	// The cached Response from VEVO API
	Response VideoJSON
}

cacheStore := map[string][CachedVideoJson]

// Takes given VideoJSON |input| and saves it into our in-memory datastore
// using the key |cachekey|. This is so that the response can quickly
// be looked up in the future. |r| is the Http request for the current session.
func CacheVideoJsonResponse(r *http.Request, input VideoJSON, cacheKey string) (err error) {
	cacheKey = strings.ToLower(cacheKey)
	cachedData := CachedVideoJson{
		VideoId:  cacheKey,
		Response: input,
	}

	// Try to determine if cache entry already exists
	cachedValue, ok = cacheStore[cacheKey]
	if !ok {
		// Cache entry does not exist, so we should store it.
		cacheStore[cacheKey] = cachedData
	}
	return nil
}

// Attempts to find a cached response given |videoid|. Note that |videoid| should
// be the primary key we use to store the cached responses. Returns true iff a valid
// response was found.
func GetCachedVideoResponse(r *http.Request, videoid string) (found bool, response *VideoJSON, err error) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery(kTableNameCachedVideo).Filter("VideoId=", videoid)
	// Make a array slice that's initially empty but has capacity 5.
	videos := make([]CachedVideoJson, 0, 5)
	if _, e := q.GetAll(c, &videos); e != nil {
		// failed to execute query
		err = fmt.Errorf("Error:%v\nFailed to retrieve cached response from datastore", e)
		return false, nil, err
	}
	if len(videos) > 0 {
		// Found at least a viable cached video response
		return true, &videos[0].Response, nil
	} else {
		// No cached video response found.
		return false, nil, err
	}

}
