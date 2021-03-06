package db

import "log"

// GetOrganisationID gets organisation user belongs to
func (p *PostgresDb) GetOrganisationID(userID string) (orgID string, err error) {
	query :=
		`
			SELECT org_id 
			FROM organisations
			WHERE owner_id = $1
		`
	err = p.dbConn.Get(&orgID, query, userID)
	if err != nil {
		log.Printf("Error getting organisation for userID - %v: %v", userID, err.Error())
		return
	}
	return
}
