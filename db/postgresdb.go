package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used to specify postgres driver
)

// PostgresDb stores a connection to a postgres db
// also implements Db interface
type PostgresDb struct {
	dbConn *sqlx.DB
}

//InitDb creates a table in postgres using the configuration provided
func InitDb(connStr string) (*PostgresDb, error) {
	dbConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	p := &PostgresDb{
		dbConn: dbConn,
	}
	err = p.dbConn.Ping()
	if err != nil {
		return nil, err
	}
	err = p.createTable()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PostgresDb) createTable() (err error) {
	users :=
		`
			CREATE TABLE IF NOT EXISTS user_accounts (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				user_id 	TEXT UNIQUE,
				id_type		TEXT 
			)
		`

	organisations :=
		`
			CREATE TABLE IF NOT EXISTS organisations (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				org_id 		TEXT UNIQUE,
				owner_id	TEXT REFERENCES user_accounts(user_id)
			)
		`

	customers :=
		`
			CREATE TABLE IF NOT EXISTS customers (
				id 				BIGSERIAL NOT NULL PRIMARY KEY,
				customer_id 	TEXT NOT NULL,
				org_id 			TEXT NOT NULL,
				name			TEXT NOT NULL,
				email	  		TEXT,
				phone_no 		TEXT,
				gstin			TEXT, 
				CONSTRAINT fk_org_id FOREIGN KEY(org_id) REFERENCES organisations(org_id),
				CONSTRAINT unique_customer_org UNIQUE (customer_id, org_id)
			)	
		`

	items :=
		`
			CREATE TABLE IF NOT EXISTS items (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				item_id 	TEXT,
				org_id  	TEXT,
				name 		TEXT,
				type 		TEXT,
				hsn 		INTEGER,
				sac 		INTEGER,
				gst 		REAL,
				igst 		REAL,
				CONSTRAINT fk_org_id FOREIGN KEY(org_id) REFERENCES organisations(org_id),
				CONSTRAINT unique_item_org UNIQUE (item_id, org_id)
			)
		`

	_, err = p.dbConn.Exec(users)
	if err != nil {
		return
	}

	_, err = p.dbConn.Exec(organisations)
	if err != nil {
		return
	}

	_, err = p.dbConn.Exec(customers)
	if err != nil {
		return
	}

	_, err = p.dbConn.Exec(items)
	if err != nil {
		return
	}

	return
}
