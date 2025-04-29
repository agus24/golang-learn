package helpers

import (
	"database/sql"
	"fmt"
	"golang_gin/app/databases/model"
	"golang_gin/tests/seeders"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
)

type HttpRequestBuilder struct {
	Method  string
	URL     string
	Body    map[string]any
	BaseUrl string
	Token   *string
	db      *sql.DB
}

func NewHttpRequestBuilder(baseUrl string, db *sql.DB) *HttpRequestBuilder {
	return &HttpRequestBuilder{
		BaseUrl: baseUrl,
		db:      db,
	}
}

func (self *HttpRequestBuilder) SetMethod(method string) *HttpRequestBuilder {
	self.Method = method
	return self
}

func (self *HttpRequestBuilder) SetURL(url string) *HttpRequestBuilder {
	self.URL = url
	return self
}

func (self *HttpRequestBuilder) SetBody(body map[string]any) *HttpRequestBuilder {
	self.Body = body
	return self
}

func (self *HttpRequestBuilder) SetBaseUrl(baseUrl string) *HttpRequestBuilder {
	self.BaseUrl = baseUrl
	return self
}

func (self *HttpRequestBuilder) SetToken(user *model.Users) *HttpRequestBuilder {
	if user == nil {
		user = seeders.SeedUser(self.db, "user1", "password", "user 1")
	}

	var token string = GetToken(user)

	self.Token = &token
	return self
}

func (self *HttpRequestBuilder) MapToQueryString(m map[string]any) string {
	values := url.Values{}
	for key, value := range m {
		values.Set(key, self.toString(value))
	}
	return values.Encode()
}

func (self *HttpRequestBuilder) toString(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case int, int64, float64, float32:
		return fmt.Sprintf("%v", t)
	case bool:
		return fmt.Sprintf("%t", t)
	default:
		return ""
	}
}

func (self *HttpRequestBuilder) Build() (*http.Request, error) {
	var url string = fmt.Sprintf("%s"+self.URL, self.BaseUrl)
	var body io.Reader
	var err error

	if self.Method == "GET" {
		queryString := self.MapToQueryString(self.Body)
		url = fmt.Sprintf("%s?%s", url, queryString)
	} else {
		body, err = PrepareBody(self.Body)
		if err != nil {
			log.Fatal(err, debug.Stack())
		}
	}

	req, err := http.NewRequest(self.Method, url, body)

	if err != nil {
		return nil, err
	}

	if self.Token != nil {
		req.Header.Set("Authorization", "Bearer "+*self.Token)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (self *HttpRequestBuilder) BuildAndRun(method string, url string, body map[string]any) (*http.Response, error) {
	self.SetMethod(method)
	self.SetURL(url)
	if body != nil {
		self.SetBody(body)
	}

	req, err := self.Build()
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
