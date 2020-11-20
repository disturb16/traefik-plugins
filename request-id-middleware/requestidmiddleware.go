// Package example a example plugin.
package example

import (
	"context"
	"net/http"

	"github.com/dchest/uniuri"
)

var defaultHeader string = "X-Request-Id"

// Config the plugin configuration.
type Config struct {
	HeaderName string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

// ReqMiddleware a plugin.
type ReqMiddleware struct {
	next       http.Handler
	name       string
	headerName string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &ReqMiddleware{
		next:       next,
		name:       name,
		headerName: config.HeaderName,
	}, nil
}

func (rm *ReqMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqId := uniuri.New()

	r.Header.Set(rm.headerName, reqId)
	w.Header().Set(rm.headerName, reqId)

	rm.next.ServeHTTP(w, r)
}
