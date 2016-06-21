package provider

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//Response provider response
type Response struct {
	Status     int
	Headers    http.Header
	contentSet bool
	httpContent
}

// NewResponse returns response without any body content
func NewResponse(status int, headers http.Header) *Response {
	return &Response{
		Status:  status,
		Headers: headers,
	}
}

//NewJSONResponse creates new response with body as json content
func NewJSONResponse(status int, headers http.Header) *Response {
	return &Response{
		Status:      status,
		Headers:     headers,
		httpContent: &jsonContent{},
	}
}

// NewPlainTextResponse creates new response with body as plain text content
func NewPlainTextResponse(status int, headers http.Header) *Response {
	return &Response{
		Status:      status,
		Headers:     headers,
		httpContent: &plainTextContent{},
	}
}

//ResetContent emoves an existing contentß
func (p *Response) ResetContent() {
	p.httpContent = nil
}

//MarshalJSON custom json marshaling
func (p *Response) MarshalJSON() ([]byte, error) {
	obj := map[string]interface{}{"status": p.Status}

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
func (p *Response) UnmarshalJSON(b []byte) error {
	var obj map[string]interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	r := Response{}
	if body, ok := obj["body"]; ok {
		if err := r.SetBody(body); err != nil {
			return err
		}
	}

	if val, ok := obj["status"]; ok {
		//default number deserialised as float64
		if status, ok := val.(float64); ok {
			r.Status = int(status)
		} else {
			return errors.New("Could not unmarshal response, status value is either nil or not a int")
		}
	}

	if headers, ok := obj["headers"].(map[string]interface{}); ok {
		r.Headers = make(http.Header)
		for key, val := range headers {
			if str, ok := val.(string); ok {
				r.Headers[key] = splitHeaderKeyValues(str)
			}
		}
	}

	*p = Response(r)
	return nil
}

// CreateResponseFromHTTPResponse creates response from http.Response
func CreateResponseFromHTTPResponse(httpResp *http.Response) (*Response, error) {
	resp := NewResponse(httpResp.StatusCode, httpResp.Header)

	if httpResp.Body != nil {
		data, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}
		if len(data) > 0 {
			if strings.Contains(httpResp.Header.Get("Content-Type"), "text/plain") {
				if err = resp.SetBody(string(data)); err != nil {
					return nil, err
				}
			} else {
				var body interface{}
				if err = json.Unmarshal(data, &body); err != nil {
					return nil, err
				}
				if err = resp.SetBody(body); err != nil {
					return nil, err
				}
			}
		}
	}
	return resp, nil
}

// HasContentBeenExplictlySet returns true if the user choose to set the body of the request.
func (p *Response) HasContentBeenExplictlySet() bool {
	return p.contentSet
}

// HasContent returns true when request content has been set
func (p *Response) HasContent() bool {
	return p.httpContent != nil
}

// GetData returns bytes from the content
func (p *Response) GetData() ([]byte, error) {
	if p.HasContent() {
		return p.httpContent.GetData()
	}
	return nil, nil
}

// GetBody returns the content
func (p *Response) GetBody() interface{} {
	if p.HasContent() {
		return p.httpContent.GetBody()
	}
	return nil
}

// SetBody sets the body of the request
func (p *Response) SetBody(body interface{}) error {
	p.contentSet = true
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
	return nil
}
