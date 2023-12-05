package tracer_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	tracer "github.com/laurentpoirierfr/traefik-tracer"
)

func TestTracer(t *testing.T) {
	cfg := tracer.CreateConfig()
	cfg.Params["Key"] = "Value"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := tracer.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("header-1", "value1")
	req.Header.Add("header-2", "value2")

	handler.ServeHTTP(recorder, req)

}
