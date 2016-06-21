package comparers

import (
	"strings"

	"github.com/SEEK-Jobs/pact-go/diff"
)

func headerMatches(expected, actual map[string][]string) (bool, diff.Differences) {
	if expected == nil {
		return true, nil
	}

	normalisedExpected := make(map[string][]string, len(expected))
	for key, val := range expected {
		normalisedExpected[strings.ToLower(key)] = val
	}
	var normalisedActual map[string][]string
	if actual != nil {
		normalisedActual = make(map[string][]string)
		for key, val := range actual {
			normalisedActual[strings.ToLower(key)] = val
		}
	}

	return diff.DeepDiff(normalisedExpected, normalisedActual, &diff.DiffConfig{AllowUnexpectedKeys: true, RootPath: "[\"header\"]"})
}
