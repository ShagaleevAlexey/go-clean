package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
	"net/http"
)

type WebRequest struct {
	*http.Request
}

func NewWebRequest(request *http.Request) *WebRequest {
	return &WebRequest{Request: request}
}

func (r *WebRequest) DecodeBody(v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != io.EOF && err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

func (r *WebRequest) HeaderValue(k string) string {
	return r.Header.Get(k)
}

func (r *WebRequest) URLParamValue(k string) string {
	return chi.URLParam(r.Request, k)
}
