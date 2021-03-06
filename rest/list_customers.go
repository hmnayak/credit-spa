package rest

import (
	"log"
	"net/http"

	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// ListCustomers is a handler to get all customers of an organisation
func ListCustomers(mdl *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}
		customers, err := mdl.Db.GetAllCustomers(orgID.(string))

		if err != nil {
			ui.RespondError(w, http.StatusInternalServerError, "")
		}

		res := ui.CreateResponse(http.StatusOK, "", customers)
		ui.Respond(w, res, "")
		return
	})
}
