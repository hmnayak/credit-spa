package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// UpsertCustomer processes PUT requests to upsert customers to db
func UpsertCustomer(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var customer model.Customer
		json.NewDecoder(r.Body).Decode(&customer) // TODO: error handling

		if orgID := r.Context().Value(contextkeys.OrgID); orgID != nil {
			customer.OrganisationID = orgID.(string)
		} else {
			log.Printf("Error getting orgID from context for customerID: %v", customer.CustomerID)
			return
		}

		var isNewCustomer bool
		if len(customer.CustomerID) == 0 {
			isNewCustomer = true
			assignCustomerID(db, &customer)
		}

		err := db.UpsertCustomer(customer)
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

func assignCustomerID(db model.Db, customer *model.Customer) (err error) {
	latestCustomerID, err := db.GetLatestCustomerID(customer.OrganisationID)
	if err != nil {
		return
	}
	customer.CustomerID, err = createNewCustomerID(latestCustomerID)
	if err != nil {
		return
	}
	return
}

func createNewCustomerID(latestCustomerID string) (newID string, err error) {
	if len(latestCustomerID) == 0 {
		newID = "CUST0001"
		return
	}

	latestIDParts := strings.Split(latestCustomerID, "CUST")
	latestIDNum, err := strconv.Atoi(latestIDParts[1])
	if err != nil {
		return
	}
	newID = fmt.Sprintf("CUST%04d", latestIDNum+1)
	return
}
