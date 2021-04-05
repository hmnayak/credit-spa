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

func AuthMiddleware(authClient *auth.Client, db model.Db) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			t, err := authClient.VerifyIDToken(context.Background(), token)
			if err != nil {
				ui.RespondError(w, http.StatusUnauthorized, fmt.Sprintf("Auth not Ok: %v", err.Error()))
				return
			}

			// unique identifier for users depends on sign-up method
			// email authentication has a name associated with user
			// with oauth it is auto-generated ID
			var userID, IDType string
			if username, found := t.Claims["name"]; found {
				userID = fmt.Sprintf("%v", username)
				IDType = "user_name"
			} else {
				userID = t.UID
				IDType = "user_tid"
			}

			var orgID string
			if exists, _ := db.DoesUserExist(userID); !exists {
				orgID, _ = db.CreateUser(userID, IDType)
			} else {
				orgID, _ = db.GetOrganisationID(userID)
			}
			ctx := context.WithValue(req.Context(), contextkeys.OrgID, orgID)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
