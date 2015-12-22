package mzika_test

// Integration test that is meant to verify whether logic in 
// mzika package is still applicable/working properly. The test
// cases here make actual network requests and so are suspecitible
// to flaky failures.

import (
"strings"
"testing"
"github.com/jervisfm/mzika"
)

const (
	vid = "uscmv1500002"
)

func TestGetVideoUrl(t *testing.T) {
	url, err := mzika.GetVideoUrl(vid)
	if err != nil {
		// Test failed
		t.Errorf("Failed to retrieve url for video with id '%s'. Error ", vid, err)
	}
	url = strings.ToLower(url)
	if len(url) <= 0 || !strings.HasPrefix(url, "http") {
		t.Errorf("Url obtained is invalid (either empty/not http url). Url: '%v'", url)
	}
}

