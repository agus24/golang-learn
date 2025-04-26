package helpers

import (
	"fmt"
	"io"
	"net/http"
)

type HttpRequestBuilder struct {
	Method  string
	URL     string
	Body    io.Reader
	BaseUrl string
	Token   *string
}

func NewHttpRequestBuilder(baseUrl string) *HttpRequestBuilder {
	return &HttpRequestBuilder{
		BaseUrl: baseUrl,
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

func (self *HttpRequestBuilder) SetBody(body io.Reader) *HttpRequestBuilder {
	self.Body = body
	return self
}

func (self *HttpRequestBuilder) SetBaseUrl(baseUrl string) *HttpRequestBuilder {
	self.BaseUrl = baseUrl
	return self
}

func (self *HttpRequestBuilder) SetToken(token string) *HttpRequestBuilder {
	self.Token = &token
	return self
}

func (self *HttpRequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(self.Method, fmt.Sprintf("%s"+self.URL, self.BaseUrl), self.Body)

	if err != nil {
		return nil, err
	}

	if self.Token != nil {
		req.Header.Set("Authorization", "Bearer "+*self.Token)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (self *HttpRequestBuilder) BuildAndRun(method string, url string, body *io.Reader) (*http.Response, error) {
	self.SetMethod(method)
	self.SetURL(url)
	if body != nil {
		self.SetBody(*body)
	}

	req, err := self.Build()
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
