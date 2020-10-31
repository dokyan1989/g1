package router

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/schema"
)

// Router ...
type Router struct {
	routes []route
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.Handler
}

type ctxKey struct{}

// NewRouter ...
func NewRouter() *Router {
	return &Router{}
}

// ServeHTTP ...
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range rt.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}

// Route ...
func (rt *Router) Route(method string, pattern string, handler http.Handler) {
	rt.routes = append(rt.routes, route{
		method:  method,
		regex:   regexp.MustCompile("^" + pattern + "$"),
		handler: handler,
	})
}

// URLParams ...
func URLParams(r *http.Request) []string {
	params := r.Context().Value(ctxKey{}).([]string)
	return params
}

// BodyParams ...
func BodyParams(r *http.Request, value interface{}) error {
	return json.NewDecoder(r.Body).Decode(value)
}

// QueryParams ...
func QueryParams(r *http.Request, value interface{}) error {
	return schema.NewDecoder().Decode(value, r.URL.Query())
}
