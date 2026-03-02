package middleware

import (
	"log/slog"
	"net/http"
)

// AuthorizationHandler extracts the Authorization header and stores it in the context.
type AuthorizationHandler struct {
	handler http.Handler
}

// NewAuthorizationHandler wraps an http.Handler with authorization extraction.
func NewAuthorizationHandler(handler http.Handler) http.Handler {
	return &AuthorizationHandler{handler: handler}
}

func (h *AuthorizationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		slog.WarnContext(r.Context(), "missing authorization header")
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	ctx := WithAuthorization(r.Context(), token)
	r = r.WithContext(ctx)

	h.handler.ServeHTTP(w, r)
}
