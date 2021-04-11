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

func UpsertItem(db model.Db) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var item model.Item
		err := json.NewDecoder(r.Body).Decode(&item) // TODO: error handling

		if orgID := r.Context().Value(contextkeys.OrgID); orgID != nil {
			item.OrganisationID = orgID.(string)
		} else {
			log.Printf("Error getting orgID from context for itemID: %v", item.ItemID)
			return
		}

		var isNewItem bool
		if len(item.ItemID) == 0 {
			isNewItem = true
			assignItemID(db, &item)
		}

		err = db.UpsertItem(item)
		if err != nil {
			ui.RespondError(rw, http.StatusInternalServerError, "")
			return
		}

		if isNewItem {
			res := ui.Response{HTTPStatus: http.StatusCreated}
			ui.Respond(rw, res)
		} else {
			res := ui.Response{HTTPStatus: http.StatusOK}
			ui.Respond(rw, res)
		}
	})
}

func assignItemID(db model.Db, item *model.Item) (err error) {
	latestItemID, err := db.GetLatestItemID(item.OrganisationID)
	if err != nil {
		return
	}

	var newID string
	if len(latestItemID) == 0 {
		newID = "ITEM0001"
	} else {
		latestIDParts := strings.Split(latestItemID, "ITEM")
		latestIDNum, e := strconv.Atoi(latestIDParts[1])
		if err != nil {
			return e
		}

		newID = fmt.Sprintf("ITEM%04d", latestIDNum+1)
	}

	item.ItemID = newID
	return
}
