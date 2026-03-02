package middleware

import "net/http"

type authTransport struct {
	base http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())

	r.Header.Set("Authorization", "Bearer "+GetAuthorization(req.Context()))

	return t.base.RoundTrip(r)
}

// WithBearerAuth wraps an HTTP client to forward bearer tokens from context.
func WithBearerAuth(c *http.Client) *http.Client {
	base := c.Transport
	if base == nil {
		base = http.DefaultTransport
	}

	return &http.Client{
		Transport:     &authTransport{base: base},
		CheckRedirect: c.CheckRedirect,
		Jar:           c.Jar,
		Timeout:       c.Timeout,
	}
}
