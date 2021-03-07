package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// GetCustomer process requests to get customer details of a single customer
func GetCustomer(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		customerID := params["customerid"]

		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}

		customer, err := db.GetCustomer(customerID, orgID.(string))
		if err != nil {
			ui.RespondError(w, http.StatusInternalServerError, "")
			return
		}
		res := ui.CreateResponse(http.StatusOK, "", customer)
		ui.Respond(w, res, "")
		return
	})
}
