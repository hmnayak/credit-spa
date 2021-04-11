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

func ListItems(db model.Db) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		orgID := r.Context().Value(contextkeys.OrgID)
		if orgID == nil {
			log.Printf("No orgID in context")
			return
		}

		if pageToken := r.URL.Query().Get("page"); len(pageToken) > 0 {
			pageToken, _ := strconv.Atoi(pageToken)
			listItemsPaginated(orgID.(string), pageToken, db, rw)
		}
	})
}

func listItemsPaginated(orgID string, page int, db model.Db, rw http.ResponseWriter) {
	count, err := db.GetItemsCount(orgID)
	if err != nil {
		ui.RespondError(rw, http.StatusInternalServerError, "")
		return
	}

	if isValid := validateItemsPageToken(page, count); !isValid {
		ui.RespondError(rw, http.StatusBadRequest, "page does not exist")
		return
	}

	items, err := db.GetItemsPaginated(orgID, page, ApiConfig.ItemsPageSize)
	if err != nil {
		ui.RespondError(rw, http.StatusInternalServerError, "")
		return
	}

	nextPage := 0
	if page*ApiConfig.ItemsPageSize < count {
		nextPage = page + 1
	}

	listResponse := model.ListItemsResponse{Items: items, NextPageToken: nextPage, TotalSize: count}
	res := ui.Response{HTTPStatus: http.StatusOK, Payload: listResponse}
	ui.Respond(rw, res)
}

func validateItemsPageToken(pageToken int, itemsCount int) bool {
	return itemsCount > (pageToken-1)*ApiConfig.ItemsPageSize ||
		(itemsCount == 0 && pageToken == 1)
}
