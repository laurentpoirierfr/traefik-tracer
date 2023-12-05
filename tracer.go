package tracer

import (
	"context"
	"log"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Params map[string]string `json:"params,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Params: make(map[string]string),
	}
}

// Tracer a plugin.
type Tracer struct {
	next   http.Handler
	name   string
	params map[string]string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Tracer{
		name:   name,
		params: config.Params,
		next:   next,
	}, nil
}

func (t *Tracer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t.logHeaders(req)
	t.next.ServeHTTP(rw, req)
}

func (t *Tracer) logHeaders(req *http.Request) {
	for key, val := range req.Header {
		log.Println(key, ":", val)
	}
}
