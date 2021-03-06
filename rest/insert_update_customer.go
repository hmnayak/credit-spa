package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// UpsertCustomer processes PUT requests to upsert customers to db
func UpsertCustomer(mdl *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var customer model.Customer
		json.NewDecoder(r.Body).Decode(&customer) // TODO: error handling

		isNewCustomer := false
		if len(customer.CustomerID) == 0 {
			isNewCustomer = true
			n, err := mdl.Db.GetCustomerCount()
			if err != nil {
				return // TODO: error handling
			}
			customer.CustomerID = fmt.Sprintf("CUST%04d", n+1)
		}

		if orgID := r.Context().Value("org_id"); orgID != nil {
			customer.OrganisationID = orgID.(string)
		} else {
			log.Printf("Error getting orgID from context for customerID: %v", customer.CustomerID)
			return
		}

		err := mdl.Db.UpsertCustomer(customer)
		if err != nil {
			ui.RespondError(w, http.StatusInternalServerError, "")
			return
		}

		if isNewCustomer {
			w.WriteHeader(http.StatusCreated)
		} else {
			res := ui.CreateResponse(http.StatusOK, "", customer)
			ui.Respond(w, res, "")
		}
	})
}
