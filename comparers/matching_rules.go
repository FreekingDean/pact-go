package comparers

import ()

type Matcher struct {
	Min           int
	Max           int
	MatchType     bool
	MatchEquality bool
	MatchInteger  bool
	MatchDecimal  bool
	Regex         string
}

type Rule map[string]Matcher

type MatchingRules struct {
	Body RuleSet
}

func (mr *MatchignRules) UnmarshalJSON(b []byte) error {
	var obj map[string]interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	if body, ok := obj["body"].(map[string]interface{}); ok {
		for path, matcher := range body {
			if matcher, ok := matcher.(map
			mr.Body[path]
		}
	}
}
