package comparers

import (
	"net/url"

	"fmt"
	"github.com/SEEK-Jobs/pact-go/provider"
)

// MatchRequest compares the request and provides the match outcome
func MatchRequest(expected, actual *provider.Request) (bool, error) {
	expectedQuery, err := url.ParseQuery(expected.Query)
	if err != nil {
		return false, err
	}
	if len(expected.MatchingRules) > 0 {
		fmt.Println("%+v", expected.MatchingRules)
		panic(expected.MatchingRules)
	}
	actualQuery, err := url.ParseQuery(actual.Query)
	if err != nil {
		return false, err
	}

	if !methodMatches(expected.Method, actual.Method) {
		return false, nil
	} else if !pathMatches(expected.Path, actual.Path) {
		return false, nil
	} else if res, _ := queryMatches(expectedQuery, actualQuery); !res {
		return false, nil
	} else if res, _ := headerMatches(expected.Headers, actual.Headers); !res {
		return false, nil
	} else if res, _, err := bodyMatches(expected.GetBody(), actual.GetBody(), false, expected.BodyHasToBeSerialized()); err != nil || !res {
		return false, err
	}
	return true, nil
}
