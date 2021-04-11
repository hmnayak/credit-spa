package db

import (
	"database/sql"
	"log"

	"github.com/hmnayak/credit/model"
)

func (p *PostgresDb) UpsertItem(it model.Item) (err error) {
	query :=
		`
		INSERT INTO items (item_id, org_id, name, type, hsn, sac, gst, igst) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (item_id, org_id)
		DO UPDATE SET
			name = EXCLUDED.name,
			type = EXCLUDED.type,
			hsn = EXCLUDED.hsn,
			sac = EXCLUDED.sac, 
			gst = EXCLUDED.gst,
			igst = EXCLUDED.igst
	`

	_, err = p.dbConn.Exec(query, it.ItemID, it.OrganisationID, it.Name, it.Type, it.HSN, it.SAC, it.GST, it.IGST)
	if err != nil {
		log.Printf("Error inserting item: %v\n", err.Error())
	}
	return
}

func (p *PostgresDb) GetItemsPaginated(orgID string, pageToken int, pageSize int) (items []*model.Item, err error) {
	items = []*model.Item{}
	query :=
		`
			SELECT item_id, name, type, hsn, sac, gst, igst
			FROM items
			WHERE org_id = $1
			ORDER BY item_id ASC
			LIMIT $2 OFFSET $3
		`

	offset := 0
	if pageToken > 0 {
		offset = (pageToken - 1) * pageSize
	}

	err = p.dbConn.Select(&items, query, orgID, pageSize, offset)
	if err != nil {
		log.Printf("Error getting paginated list of items: %v", err.Error())
	}
	return
}

// GetLatestItemID gets the id assigned to last inserted item
func (p *PostgresDb) GetLatestItemID(orgID string) (itemID string, err error) {
	var nullableID sql.NullString
	query :=
		`
			SELECT MAX(item_id)
			FROM items
			WHERE org_id = $1
		`

	err = p.dbConn.Get(&nullableID, query, orgID)
	if err != nil {
		log.Printf("Error getting latest customerID: %v", err.Error())
	}

	itemID = nullableID.String
	return
}

// GetItemsCount gets the total number items
func (p *PostgresDb) GetItemsCount(orgID string) (count int, err error) {
	query :=
		`
			SELECT COUNT(*)
			FROM items
			WHERE org_id = $1
		`

	err = p.dbConn.Get(&count, query, orgID)
	if err != nil {
		log.Printf("Error getting items count: %v", err.Error())
	}
	return
}
