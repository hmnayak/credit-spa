package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hmnayak/credit/contextkeys"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// ListCustomers is a handler to get all customers of an organisation
func ListCustomers(db model.Db, pageSize int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}

		if _, ok := r.URL.Query()["page"]; ok {
			pageToken, _ := strconv.Atoi(r.URL.Query().Get("page"))

			customersCount, err := db.GetCustomersCount(orgID.(string))
			if err != nil {
				ui.RespondError(w, http.StatusInternalServerError, "")
				return
			}

			if customersCount <= (pageToken-1)*pageSize {
				ui.RespondError(w, http.StatusBadRequest, "page does not exist")
				return
			}

			customers, err := db.GetCustomersPaginated(orgID.(string), pageToken, pageSize)
			if err != nil {
				ui.RespondError(w, http.StatusInternalServerError, "")
				return
			}

			nextPageToken := 0
			if pageToken*pageSize < customersCount {
				nextPageToken = pageToken + 1
			}
			listCustomersResponse := model.ListCustomersResponse{Customers: customers, NextPageToken: nextPageToken, TotalSize: customersCount}
			res := ui.CreateResponse(http.StatusOK, listCustomersResponse)
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
