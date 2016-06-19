package comparers

import (
	_ "encoding/json"
	"net/http"
	"testing"

	"github.com/SEEK-Jobs/pact-go/provider"
)

func Test_MethodIsDifferent_WillNotMatch(t *testing.T) {
	a := provider.NewJSONRequest("GET", "", "", nil)
	b := provider.NewJSONRequest("POST", "", "", nil)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if result {
		t.Error("The request should not match")
	}
}

func Test_UrlIsDifferent_WillNotMatch(t *testing.T) {
	a := provider.NewJSONRequest("GET", "", "", nil)
	b := provider.NewJSONRequest("GET", "/", "", nil)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if result {
		t.Error("The request should not match")
	}
}

func Test_ExpectedNoBodyButActualRequestHasBody_WillMatch(t *testing.T) {
	a := provider.NewJSONRequest("GET", "/test", "", nil)
	b := provider.NewJSONRequest("GET", "/test", "", nil)
	b.SetBody(`{"name": "John"}`)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !result {
		t.Error("The request should match")
	}
}

func Test_BodyIsDifferent_WillNotMatch(t *testing.T) {
	a := provider.NewJSONRequest("GET", "/test", "", nil)
	a.SetBody(`{"name": "John", "age": 12 }`)
	b := provider.NewJSONRequest("GET", "/test", "", nil)
	b.SetBody(`{"name": "John"}`)

	result, err := MatchRequest(a, b)
	if result {
		t.Error("The request should not match")
	}

	if err != nil {
		t.Error(err)
	}

}

func Test_HeadersAreMissing_WillNotMatch(t *testing.T) {
	aHeader := make(http.Header)
	aHeader.Add("content-type", "application/json")
	a := provider.NewJSONRequest("GET", "/test", "", aHeader)
	b := provider.NewJSONRequest("GET", "/test", "", nil)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if result {
		t.Error("The request should not match")
	}
}

func Test_HeadersAreDifferent_WillNotMatch(t *testing.T) {
	aHeader := make(http.Header)
	aHeader.Add("content-type", "application/json")
	a := provider.NewJSONRequest("GET", "/test", "", aHeader)

	bHeader := make(http.Header)
	bHeader.Add("content-type", "text/plain")
	b := provider.NewJSONRequest("GET", "/test", "", bHeader)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if result {
		t.Error("The request should not match")
	}
}

func Test_AllHeadersFound_WillMatch(t *testing.T) {
	aHeader := make(http.Header)
	bHeader := make(http.Header)

	aHeader.Add("content-type", "application/json")
	bHeader.Add("content-type", "application/json")
	bHeader.Add("extra-header", "value")

	a := provider.NewJSONRequest("GET", "/test", "", aHeader)
	b := provider.NewJSONRequest("GET", "/test", "", bHeader)

	result, err := MatchRequest(a, b)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !result {
		t.Error("The request should match")
	}
}
