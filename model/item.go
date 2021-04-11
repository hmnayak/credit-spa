package model

type Item struct {
	ID             string `db:"id" json:"-"`
	ItemID         string `db:"item_id" json:"itemid"`
	OrganisationID string `db:"org_id" json:"-"`
	Name           string `db:"name" json:"name"`
	Type           string `db:"type" json:"type"`
	HSN            int    `db:"hsn" json:"hsn,string"`
	SAC            int    `db:"sac" json:"sac,string"`
	GST            int    `db:"gst" json:"gst,string"`
	IGST           int    `db:"igst" json:"igst,string"`
}

type ListItemsResponse struct {
	Items         []*Item `json:"items"`
	NextPageToken int     `json:"nextpagetoken"`
	TotalSize     int     `json:"totalsize"`
}
