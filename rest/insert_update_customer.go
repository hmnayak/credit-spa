package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

		isNewCustomer := false
		if len(customer.CustomerID) == 0 {
			isNewCustomer = true
			latestCustomerID, err := db.GetLatestCustomerID(customer.OrganisationID)
			createNewCustomerID(latestCustomerID)
			if err != nil {
				return
			}
			customer.CustomerID, err = createNewCustomerID(latestCustomerID)
			if err != nil {
				return
			}
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

func createNewCustomerID(latestCustomerID string) (newCustomerID string, err error) {
	if err != nil {
		return
	}
	rx, err := regexp.Compile("CUST")
	if err != nil {
		return
	}
	lastIDParts := rx.Split(latestCustomerID, 2)
	if err != nil {
		return
	}
	lastIDNum, err := strconv.Atoi(lastIDParts[1])
	if err != nil {
		return
	}
	newIDNum := lastIDNum + 1
	newCustomerID = fmt.Sprintf("CUST%04d", newIDNum)
	return
}
