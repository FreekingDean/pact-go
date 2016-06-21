package comparers

import "github.com/SEEK-Jobs/pact-go/diff"

func bodyMatches(expected, actual interface{}, allowUnexpectedKeys bool, expectedBody bool) (bool, diff.Differences, error) {
	if expected == nil && !expectedBody {
		return true, nil, nil
	}
	if result, diffs := diff.DeepDiff(expected, actual, &diff.DiffConfig{AllowUnexpectedKeys: allowUnexpectedKeys, RootPath: "[\"body\"]"}); result {
		return result, nil, nil
	} else {
		return result, diffs, nil
	}
}
