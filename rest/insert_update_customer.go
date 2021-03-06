package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hmnayak/credit/model"
)

// UpsertCustomer processes PUT requests to upsert customers to db
func UpsertCustomer(mdl *model.Model) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var customer model.Customer
		json.NewDecoder(req.Body).Decode(&customer) // TODO: error handling

		if len(customer.CustomerID) == 0 {
			n, err := mdl.Db.GetCustomerCount()
			if err != nil {
				return // TODO: error handling
			}
			customer.CustomerID = fmt.Sprintf("CUST%04d", n+1)
		}

		if orgID := req.Context().Value("org_id"); orgID != nil {
			customer.OrganisationID = orgID.(string)
		} else {
			log.Printf("Error getting orgID from context for customerID: %v", customer.CustomerID)
			return
		}

		mdl.Db.UpsertCustomer(customer)
	})
}
