package api

import (
	"JetChatClientGo/utils"
	"bytes"
	"encoding/json"
	"net/http"
)

var DefaultHeaders = map[string]string{
	"User-Agent": "JetChatClientGo",
	"Connection": "keep-alive",
}

type Request[T any] struct {
	client    *http.Client
	Body      any
	BodyType  string
	Cookies   []*http.Cookie
	Headers   map[string]string
	Method    string
	URL       string
	URLParams map[string]string
}

func (r *Request[T]) AddURLParam(key string, value string) {
	r.URLParams[key] = value
}

func (r *Request[T]) AddHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *Request[T]) SetBody(body any) {
	r.Body = body
}

func (r *Request[T]) Json() *Request[T] {
	r.AddHeader("Content-Type", "application/json")
	return r
}

func (r *Request[T]) Session(rs *http.Request) *Request[T] {
	sessionCookie, err := rs.Cookie(utils.ConnectionCookie)
	if err != nil {
		return r
	}

	r.Cookies = append(r.Cookies, sessionCookie)
	return r
}

func (r *Request[T]) Send() (*http.Response, error) {
	req, response, err := r.prepareRequest()
	if err != nil {
		return response, err
	}

	resp, err := r.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

func (r *Request[T]) SendWithResponse() (*T, *http.Response, error) {
	req, response, err := r.prepareRequest()
	if err != nil {
		return nil, response, err
	}

	resp, err := r.client.Do(req)

	if err != nil {
		return nil, resp, err
	}

	defer resp.Body.Close()

	var body T
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		return nil, resp, err
	}

	return &body, resp, nil
}

func (r *Request[T]) prepareRequest() (*http.Request, *http.Response, error) {
	jsonBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, nil, err
	}

	for key, value := range DefaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	for key, value := range r.URLParams {
		req.URL.Query().Add(key, value)
	}

	for _, cookie := range r.Cookies {
		req.AddCookie(cookie)
	}

	return req, nil, err
}

func NewRequest[T any](route string) *Request[T] {
	return &Request[T]{
		client:    &http.Client{},
		URL:       ApiURL + route,
		URLParams: make(map[string]string),
		Headers:   make(map[string]string),
		Method:    "GET",
	}
}
