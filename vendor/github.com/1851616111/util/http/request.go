package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ContentType string

var ContentType_JSON ContentType = "application/json"
var ContentType_FORM ContentType = "application/x-www-form-urlencoded"

func NewRequest(spec *HttpSpec) (*http.Request, error) {
	var req *http.Request
	var body io.Reader = nil
	var err error
	switch spec.Method {
	case "POST":
		if spec.BodyParams != nil && len(*spec.BodyParams) > 0 {
			switch spec.ContentType {
			case ContentType_JSON:
				buf := &bytes.Buffer{}
				encoder := json.NewEncoder(buf)
				encoder.SetEscapeHTML(false)
				if err := encoder.Encode(spec.BodyParams); err != nil {
					return nil, err
				}
				body = io.Reader(buf)
			case ContentType_FORM:
				v := url.Values{}
				for key, value := range *spec.BodyParams {
					if s, ok := value.(string); ok {
						v.Add(key, s)
					}
				}
				body = ioutil.NopCloser(strings.NewReader(v.Encode()))
			}
		}
	case "GET":
	}

	urlStr, urlParam := spec.URL, "?"
	if spec.URLParams != nil {
		for idx, param := range *spec.URLParams {
			urlParam += fmt.Sprintf("%s=%s", param.Name, param.Value)
			if idx != len(*spec.URLParams)-1 {
				urlParam += "&"
			}
		}
		if urlParam != "?" {
			urlStr += urlParam
		}
	}

	if req, err = http.NewRequest(spec.Method, urlStr, body); err != nil {
		return nil, err
	}

	if len(spec.Header) > 0 {
		for k, v := range spec.Header {
			req.Header.Set(k, v)
		}
	}

	if spec.BasicAuth != nil {
		req.SetBasicAuth(spec.BasicAuth.User, spec.BasicAuth.Password)
	}

	req.Header.Set("Content-Type", string(spec.ContentType))

	return req, nil
}

type HttpSpec struct {
	URL         string            `json:"url"`
	Method      string            `json:"method"`
	ContentType ContentType       `json:"content_type"`
	URLParams   *Params           `json:"url_params"`
	BodyParams  *Body             `json:"body_params"`
	Header      map[string]string `json:"header"`
	BasicAuth   *BasicAuth        `json:"basicauth"`
}

type BasicAuth struct {
	User     string
	Password string
}

type Body map[string]interface{}

func NewBody() *Body {
	return &Body{}
}

func (p *Body) Add(key string, value interface{}) *Body {
	map[string]interface{}(*p)[key] = value
	return p
}

type Params []Param

type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewParams() *Params {
	return &Params{}
}

func (p *Params) Add(key, value string) *Params {
	(*p) = append((*p), Param{key, value})
	return p
}

func (p *Params) Set(key, value string) *Params {
	for idx, param := range *p {
		if param.Name == key {
			(*p)[idx].Value = value
		}
	}
	return p
}

func (p *Params) Get(key string) (bool, string) {
	for _, param := range *p {
		if param.Name == key {
			return true, param.Value
		}
	}

	return false, ""
}

func (p *Params) Rename(old, new string) {
	for id, param := range *p {
		if param.Name == old {
			(*p)[id].Name = new
			return
		}
	}
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func (p *Params) Len() int {
	slice := ([]Param)(*p)
	return len(slice)
}

func (p *Params) Less(i, j int) bool {
	slice := ([]Param)(*p)
	return slice[i].Name < slice[j].Name
}

func (p *Params) Swap(i, j int) {
	slice := ([]Param)(*p)
	slice[i], slice[j] = slice[j], slice[i]
}
