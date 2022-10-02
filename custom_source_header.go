package custom_source_header

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	HeaderName string `yaml:"headerName"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// CustomSourceHeader a CustomSourceHeader plugin.
type CustomSourceHeader struct {
	next   http.Handler
	name   string
	config *Config
}

// New created a new CustomSourceHeader plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.HeaderName == "" {
		return nil, fmt.Errorf("headerName cannot be empty")
	}

	return &CustomSourceHeader{
		config: config,
		next:   next,
		name:   name,
	}, nil
}

func (a *CustomSourceHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	remoteAddr := req.RemoteAddr
	host, _, splitError := net.SplitHostPort(remoteAddr)
	if splitError != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set(a.config.HeaderName, host)
	a.next.ServeHTTP(rw, req)
}
