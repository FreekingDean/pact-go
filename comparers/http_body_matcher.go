package comparers

import (
	"encoding/json"
	"io"

	"github.com/SEEK-Jobs/pact-go/diff"
)

func bodyMatches(expected, actual io.Reader, allowUnexpectedKeys bool) (bool, diff.Differences, error) {
	if expected == nil {
		return true, nil, nil
	}
	var e, a interface{}
	decoder := json.NewDecoder(expected)
	err := decoder.Decode(&e)
	if err != nil {
		return false, nil, err
	}

	if actual != nil {
		decoder = json.NewDecoder(actual)
		err = decoder.Decode(&a)
		if err != nil {
			return false, nil, err
		}
	}

	if result, diffs := diff.DeepDiff(e, a, &diff.DiffConfig{AllowUnexpectedKeys: allowUnexpectedKeys, RootPath: "[\"body\"]"}); result {
		return result, nil, nil
	} else {
		return result, diffs, nil
	}
}

func bodyMatchesTemp(expected, actual interface{}, allowUnexpectedKeys bool) (bool, diff.Differences, error) {
	if expected == nil {
		return true, nil, nil
	}
	if result, diffs := diff.DeepDiff(expected, actual, &diff.DiffConfig{AllowUnexpectedKeys: allowUnexpectedKeys, RootPath: "[\"body\"]"}); result {
		return result, nil, nil
	} else {
		return result, diffs, nil
	}
}
