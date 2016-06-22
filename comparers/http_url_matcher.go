package comparers

import (
	"net/url"

	"github.com/SEEK-Jobs/pact-go/diff"
)

func pathMatches(expected, actual string) bool {
	if expected != actual {
		return false
	}
	return true
}

func queryMatches(expected, actual url.Values) (bool, diff.Differences) {
	return diff.DeepDiff(expected, actual, &diff.DiffConfig{AllowUnexpectedKeys: false})
}
