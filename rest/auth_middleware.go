package rest

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/hmnayak/credit/ui"
)

func AuthMiddleware(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			_, err := authClient.VerifyIDToken(context.Background(), token)
			if err != nil {
				origin := req.Header.Get("Origin")
				var response ui.Response
				response = ui.CreateResponse(http.StatusUnauthorized, "Auth not Ok", nil)
				ui.Respond(w, response, origin)
				return
			} else {
				next.ServeHTTP(w, req)
			}
		})
	}
}
