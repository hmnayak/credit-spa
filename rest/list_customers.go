package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hmnayak/credit/config"
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

		if _, ok := r.URL.Query()["page"]; ok {
			pageToken, _ := strconv.Atoi(r.URL.Query().Get("page"))

			nCustomers, err := db.GetCustomersCount(orgID.(string))
			if err != nil {
				ui.RespondError(w, http.StatusInternalServerError, "")
				return
			}

			if nCustomers > 0 && nCustomers <= (pageToken-1)*config.ApiConfig.CustomersPageSize {
				ui.RespondError(w, http.StatusBadRequest, "page does not exist")
				return
			}

			customers, err := db.GetCustomersPaginated(orgID.(string), pageToken, config.ApiConfig.CustomersPageSize)
			if err != nil {
				ui.RespondError(w, http.StatusInternalServerError, "")
				return
			}

			nextPage := 0
			if pageToken*config.ApiConfig.CustomersPageSize < nCustomers {
				nextPage = pageToken + 1
			}
			listResponse := model.ListCustomersResponse{Customers: customers, NextPageToken: nextPage, TotalSize: nCustomers}
			res := ui.CreateResponse(http.StatusOK, listResponse)
			ui.Respond(w, res)
		} else {
			customers, err := db.GetAllCustomers(orgID.(string))
			if err != nil {
				ui.RespondError(w, http.StatusInternalServerError, "")
			}
			res := ui.CreateResponse(http.StatusOK, customers)
			ui.Respond(w, res)
		}
	})
}
