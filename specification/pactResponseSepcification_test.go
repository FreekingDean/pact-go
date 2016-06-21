package specification

import (
	"io/ioutil"
	"testing"

	"encoding/json"

	"github.com/SEEK-Jobs/pact-go/comparers"
	"github.com/SEEK-Jobs/pact-go/provider"
)

type ResponseTestCase struct {
	Match    bool               `json:"match"`
	Comment  string             `json:"comment"`
	Expected *provider.Response `json:"expected"`
	Actual   *provider.Response `json:"actual"`
}

func TestResponseBodySpecificaion(t *testing.T) {
	searchDir := "./pact-specification/testcases/response/body/"

	fileList, err := getFileNamesFromFolder(searchDir)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, fileName := range fileList {
		testResponseCase(t, fileName)
	}
}

func TestResponseHeaderSpecificaion(t *testing.T) {
	searchDir := "./pact-specification/testcases/response/headers/"

	fileList, err := getFileNamesFromFolder(searchDir)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, fileName := range fileList {
		testResponseCase(t, fileName)
	}
}

func TestResponseStatusSpecificaion(t *testing.T) {
	searchDir := "./pact-specification/testcases/Response/status/"

	fileList, err := getFileNamesFromFolder(searchDir)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, fileName := range fileList {
		testResponseCase(t, fileName)
	}
}

func testResponseCase(t *testing.T, fileName string) {
	data, err := ioutil.ReadFile("./" + fileName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tc := &ResponseTestCase{}
	if err := json.Unmarshal(data, tc); err != nil {
		t.Error(err)
		t.FailNow()
	}

	diffs, err := comparers.MatchResponse(tc.Expected, tc.Actual)
	match := (len(diffs) == 0)
	if err != nil {
		t.Error(err)
	} else if match != tc.Match {
		t.Error(tc.Comment)
	}
}
