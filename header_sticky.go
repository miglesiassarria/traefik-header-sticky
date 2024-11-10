package traefik_header_sticky

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	HeaderName string `json:"headerName,omitempty"`
	CookieName string `json:"cookieName,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderName: "X-Custom-Header",
		CookieName: "sticky-header",
	}
}

// HeaderSticky a plugin.
type HeaderSticky struct {
	next       http.Handler
	headerName string
	cookieName string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HeaderSticky{
		next:       next,
		headerName: config.HeaderName,
		cookieName: config.CookieName,
	}, nil
}

func (h *HeaderSticky) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	headerValue := req.Header.Get(h.headerName)
	if headerValue != "" {
		cookie := &http.Cookie{
			Name:  h.cookieName,
			Value: headerValue,
			Path:  "/",
		}
		http.SetCookie(rw, cookie)
	}
	h.next.ServeHTTP(rw, req)
}
