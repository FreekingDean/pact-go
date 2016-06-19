package specification

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"encoding/json"

	"github.com/SEEK-Jobs/pact-go/comparers"
	"github.com/SEEK-Jobs/pact-go/consumer"
	"github.com/SEEK-Jobs/pact-go/provider"
)

type RequestTestCase struct {
	Match    bool              `json:"match"`
	Comment  string            `json:"comment"`
	Expected *provider.Request `json:"expected"`
	Actual   *provider.Request `json:"actual"`
}

func convertToHTTPRequest(r *provider.Request) (*http.Request, error) {
	i := &consumer.Interaction{Request: r}
	return i.ToHTTPRequest("http://localhost")
}

func getFileNamesFromFolder(folderPath string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(folderPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}

func TestRequestBodySpecificaion(t *testing.T) {
	searchDir := "./pact-specification/testcases/request/body/"

	fileList, err := getFileNamesFromFolder(searchDir)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, fileName := range fileList {
		data, err := ioutil.ReadFile("./" + fileName)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		tc := &RequestTestCase{}
		if err := json.Unmarshal(data, tc); err != nil {
			t.Error(err)
			t.FailNow()
		}

		result, err := comparers.MatchRequest(tc.Expected, tc.Actual)
		if err != nil {
			t.Log(tc.Comment)
			t.Error(err)
		} else if result != tc.Match {
			t.Errorf("Expected Match: %v Actual Match: %v", tc.Match, result)
			t.Error(tc.Comment)
			t.Logf("Exp: %v \r\nAct: %v", tc.Expected.GetBody(), tc.Actual.GetBody())
		}
	}
}
