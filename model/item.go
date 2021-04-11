package model

type Item struct {
	ID             string `db:"id" json:"-"`
	ItemID         string `db:"item_id" json:"itemid"`
	OrganisationID string `db:"org_id" json:"-"`
	Name           string `db:"name" json:"name"`
	Type           string `db:"type" json:"type"`
	HSN            int    `db:"hsn" json:"hsn"`
	SAC            int    `db:"sac" json:"sac"`
	GST            int    `db:"gst" json:"gst"`
	IGST           int    `db:"igst" json:"igst"`
}

type ListItemsResponse struct {
	Items         []*Item `json:"items"`
	NextPageToken int     `json:"nextpagetoken"`
	TotalSize     int     `json:"totalsize"`
}
