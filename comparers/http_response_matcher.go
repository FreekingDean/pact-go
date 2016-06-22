package comparers

import (
	"github.com/SEEK-Jobs/pact-go/diff"
	"github.com/SEEK-Jobs/pact-go/provider"
)

func MatchResponse(expected, actual *provider.Response) (diff.Differences, error) {
	diffs := make(diff.Differences, 0)

	if res, sDiff := diff.DeepDiff(expected.Status, actual.Status,
		&diff.DiffConfig{AllowUnexpectedKeys: true, RootPath: "[\"status\"]"}); !res {
		diffs = append(diffs, sDiff...)
	} else if res, hDiff := headerMatches(expected.Headers, actual.Headers); !res {
		diffs = append(diffs, hDiff...)
	} else if res, bDiff, err := bodyMatches(expected.GetBody(), actual.GetBody(), true, expected.BodyHasToBeSerialized()); err != nil {
		return nil, err
	} else if !res {
		diffs = append(diffs, bDiff...)
	}

	return diffs, nil
}
