package rest

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

func AuthMiddleware(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			_, err := authClient.VerifyIDToken(context.Background(), token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				next.ServeHTTP(w, req)
			}
		})
	}
}
