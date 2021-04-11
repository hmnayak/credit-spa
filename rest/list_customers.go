package rest

import (
	"log"
	"net/http"
	"strconv"

	. "github.com/hmnayak/credit/config"
	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// ListCustomers is a handler to get all customers of an organisation
func ListCustomers(db model.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}

		if pageToken := r.URL.Query().Get("page"); len(pageToken) > 0 {
			pageToken, _ := strconv.Atoi(pageToken)
			listCustomersPaginated(orgID.(string), pageToken, db, w)
		} else {
			listAllCustomers(orgID.(string), db, w)
		}
	})
}

func listAllCustomers(orgID string, db model.Db, w http.ResponseWriter) {
	customers, err := db.GetAllCustomers(orgID)
	if err != nil {
		ui.RespondError(w, http.StatusInternalServerError, "")
	}
	res := ui.Response{HTTPStatus: http.StatusOK, Payload: customers}
	ui.Respond(w, res)
}

func listCustomersPaginated(orgID string, page int, db model.Db, w http.ResponseWriter) {
	count, err := db.GetCustomersCount(orgID)
	if err != nil {
		ui.RespondError(w, http.StatusInternalServerError, "")
		return
	}

	if isValid := validatePageToken(page, count); !isValid {
		ui.RespondError(w, http.StatusBadRequest, "page does not exist")
		return
	}

	customers, err := db.GetCustomersPaginated(orgID, page)
	if err != nil {
		ui.RespondError(w, http.StatusInternalServerError, "")
		return
	}

	nextPage := 0
	if page*ApiConfig.CustomersPageSize < count {
		nextPage = page + 1
	}

	res := model.ListCustomersResponse{Customers: customers, NextPageToken: nextPage, TotalSize: count}
	r := ui.Response{HTTPStatus: http.StatusOK, Payload: res}
	ui.Respond(w, r)
}

func validatePageToken(pageToken int, customersCount int) bool {
	return customersCount > (pageToken-1)*ApiConfig.CustomersPageSize ||
		(customersCount == 0 && pageToken == 1)
}
