package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//Request provider request
type Request struct {
	Method     string
	Path       string
	Query      string
	Headers    http.Header
	contentSet bool
	httpContent
}

//NewRequest creates new http request
func NewRequest(method, path, query string, headers http.Header) *Request {
	if method == "" {
		//return error
	}

	if path == "" {
		//return error
	}

	return &Request{
		Method:  method,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}

//NewJSONRequest creates new http request with content body as json
func NewJSONRequest(method, path, query string, headers http.Header) *Request {
	req := NewRequest(method, path, query, headers)
	req.httpContent = &jsonContent{}
	return req
}

//NewPlainTextRequest creates new http request with content body as json
func NewPlainTextRequest(method, path, query string, headers http.Header) *Request {
	req := NewRequest(method, path, query, headers)
	req.httpContent = &plainTextContent{}
	return req
}

//ResetContent removes an existing http content
func (p *Request) ResetContent() {
	p.httpContent = nil
}

//MarshalJSON custom json marshaling
func (p *Request) MarshalJSON() ([]byte, error) {
	obj := map[string]interface{}{
		"method": p.Method,
		"path":   p.Path,
	}

	if p.Query != "" {
		obj["query"] = p.Query
	}

	if p.Headers != nil {
		obj["headers"] = joinHeaderKeyValues(p.Headers)
	}

	if p.httpContent != nil {
		body := p.GetBody()
		if p.contentSet {
			obj["body"] = body
		}
	}

	return json.Marshal(obj)
}

//UnmarshalJSON cusotm json unmarshalling
func (p *Request) UnmarshalJSON(b []byte) error {
	var obj map[string]interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	r := Request{}

	if body, ok := obj["body"]; ok {
		r.contentSet = true
		if err := r.SetBody(body); err != nil {
			return err
		}
	}

	if method, ok := obj["method"].(string); ok {
		r.Method = method
	} else {
		return errors.New("Could not unmarshal request, method value is either nil or not a string")
	}

	if path, ok := obj["path"].(string); ok {
		r.Path = path
	} else {
		return errors.New("Could not unmarshal request, path value is either nil or not a string")
	}

	if query, ok := obj["query"].(string); ok {
		r.Query = query
	}

	if headers, ok := obj["headers"].(map[string]interface{}); ok {
		r.Headers = make(http.Header)
		for key, val := range headers {
			if str, ok := val.(string); ok {
				r.Headers[key] = splitHeaderKeyValues(str)
			}
		}
	}
	*p = Request(r)
	return nil
}

// CreateRequestFromHTTPRequest creates provdier request from http request
func CreateRequestFromHTTPRequest(httpReq *http.Request) (*Request, error) {
	req := NewRequest(httpReq.Method, httpReq.URL.Path, httpReq.URL.RawQuery, httpReq.Header)
	if httpReq.Body != nil {
		data, err := ioutil.ReadAll(httpReq.Body)
		if err != nil {
			return nil, err
		}
		if len(data) > 0 {
			switch httpReq.Header.Get("Content-Type") {
			case "text/plain":
				n := bytes.IndexByte(data, 0)
				if err = req.SetBody(string(data[:n])); err != nil {
					return nil, err
				}
			default: //expecting json
				var body interface{}
				if err = json.Unmarshal(data, &body); err != nil {
					return nil, err
				}
				if err = req.SetBody(body); err != nil {
					return nil, err
				}
			}
		}
	}
	return req, nil
}

// HasContent returns true when request content has been set
func (p *Request) HasContent() bool {
	return p.httpContent != nil
}

// HasContentBeenExplictlySet returns true if the user choose to set the body of the request.
func (p *Request) HasContentBeenExplictlySet() bool {
	return p.contentSet
}

// GetData returns bytes from the content
func (p *Request) GetData() ([]byte, error) {
	if p.HasContent() {
		return p.httpContent.GetData()
	}
	return nil, nil
}

// GetBody returns the content
func (p *Request) GetBody() interface{} {
	if p.HasContent() {
		return p.httpContent.GetBody()
	}
	return nil
}

// SetBody sets the body of the request
func (p *Request) SetBody(body interface{}) error {
	if body == nil {
		return nil
	}

	if p.httpContent == nil {
		switch body.(type) {
		case string:
			p.httpContent = &plainTextContent{}
		default:
			p.httpContent = &jsonContent{}
		}
	}

	if err := p.httpContent.SetBody(body); err != nil {
		return err
	}
	p.contentSet = true
	return nil
}

func joinHeaderKeyValues(headers http.Header) map[string]string {
	if headers == nil {
		return nil
	}

	h := make(map[string]string)
	for header, val := range headers {
		h[header] = strings.Join(val, ",")
	}
	return h
}

func splitHeaderKeyValues(val string) []string {
	splitVals := strings.Split(val, ",")
	for i := range splitVals {
		splitVals[i] = strings.TrimSpace(splitVals[i])
	}
	return splitVals
}
