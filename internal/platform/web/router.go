package web

import (
	"github.com/go-chi/chi"
	"net/http"
)

var (
	HeaderContentJson      = map[string]string{"Content-Type": "application/json"}
	HeaderContentPlainText = map[string]string{"Content-Type": "text/plain"}
)

type HandlerFunc func(*WebRequest) *WebResponse

type Router interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	Route(pattern string, fn func(r Router)) Router
	Mount(method, pattern string, handler HandlerFunc)
}

type ChiRouter struct {
	*chi.Mux
}

func NewRouter(r *chi.Mux) *ChiRouter {
	return &ChiRouter{Mux: r}
}

func (mx *ChiRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	mx.Mux.Use(middlewares...)
}

func (mx *ChiRouter) Route(pattern string, fn func(r Router)) Router {
	subRouter := NewRouter(chi.NewRouter())
	if fn != nil {
		fn(subRouter)
	}

	mx.Mux.Mount(pattern, subRouter)

	return subRouter
}

func (mx *ChiRouter) Mount(method, pattern string, handler HandlerFunc) {
	mx.MethodFunc(method, pattern, func(w http.ResponseWriter, r *http.Request) {
		webResponse := handler(NewWebRequest(r))

		if err := writeResponseResult(w, webResponse); err != nil {
			writeResponseError(w, err)
			return
		}
	})
}

func writeResponseResult(w http.ResponseWriter, resp *WebResponse) error {
	if resp.Headers != nil {
		for k, v := range resp.Headers {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.Status)

	if resp.Body != nil {
		_, err := w.Write(resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeResponseError(w http.ResponseWriter, error error) {
	if err := writeResponseResult(w, NewWebResponse(http.StatusInternalServerError, HeaderContentPlainText, []byte(error.Error()))); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}