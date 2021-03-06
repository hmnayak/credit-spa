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

func (p *PostgresDb) createTable() error {
	user :=
		`
			CREATE TABLE IF NOT EXISTS user_accounts (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				user_id 	TEXT UNIQUE,
				id_type		TEXT 
			)
		`

	organisation :=
		`
			CREATE TABLE IF NOT EXISTS organisations (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				org_id 		TEXT UNIQUE,
				owner_id	TEXT REFERENCES user_accounts(user_id)
			)
		`

	customer :=
		`
			CREATE TABLE IF NOT EXISTS customers (
				id 				BIGSERIAL NOT NULL PRIMARY KEY,
				customer_id 	TEXT NOT NULL UNIQUE,
				org_id 			TEXT NOT NULL UNIQUE,
				name			TEXT NOT NULL,
				email	  		TEXT,
				phone_no 		TEXT,
				gstin			VARCHAR(15) UNIQUE, 
				CONSTRAINT fk_org_id FOREIGN KEY(org_id) REFERENCES organisations(org_id),
				CONSTRAINT unique_customer_org UNIQUE (customer_id, org_id)
			)	
		`

	_, err := p.dbConn.Exec(user)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(organisation)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(customer)
	if err != nil {
		return err
	}

	return nil
}
