package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

type contextKey string

const orgIDKey contextKey = "org_id"

func AuthMiddleware(authClient *auth.Client, mdl *model.Model) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			t, err := authClient.VerifyIDToken(context.Background(), token)
			if err != nil {
				origin := req.Header.Get("Origin")
				var response ui.Response
				response = ui.CreateResponse(http.StatusUnauthorized, fmt.Sprintf("Auth not Ok: %v", err.Error()), nil)
				ui.Respond(w, response, origin)
				return
			}
			var userID string
			var IDType string
			if username, found := t.Claims["name"]; found {
				userID = fmt.Sprintf("%v", username)
				IDType = "user_name"
			} else {
				userID = t.UID
				IDType = "user_tid"
			}
			var orgID string
			if exists, _ := mdl.Db.DoesUserExist(userID); !exists {
				orgID, _ = mdl.Db.CreateUser(userID, IDType)
			} else {
				orgID, _ = mdl.Db.GetOrganisationID(userID)
			}
			ctx := context.WithValue(req.Context(), contextkeys.OrgID, orgID)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
