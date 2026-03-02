// Package middleware provides HTTP middleware for authorization.
package middleware

import (
	"context"
	"strings"
)

type authorizationContextKey struct{}

// WithAuthorization stores the bearer token in the context.
func WithAuthorization(ctx context.Context, token string) context.Context {
	if v, ok := strings.CutPrefix(token, "Bearer "); ok {
		token = v
	}

	return context.WithValue(ctx, authorizationContextKey{}, token)
}

// GetAuthorization retrieves the bearer token from the context.
func GetAuthorization(ctx context.Context) string {
	token, ok := ctx.Value(authorizationContextKey{}).(string)
	if !ok {
		return ""
	}

	return token
}
