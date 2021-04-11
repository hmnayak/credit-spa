package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// GetItem process requests to get details of a single item
func GetItem(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		itemID := params["itemid"]

		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}

		item, err := db.GetItem(itemID, orgID.(string))
		if err != nil {
			ui.RespondError(w, http.StatusInternalServerError, "")
			return
		}

		res := ui.Response{HTTPStatus: http.StatusOK, Payload: item}
		ui.Respond(w, res)
	})
}
