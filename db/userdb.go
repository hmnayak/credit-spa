package db

import (
	"context"
	"log"

	"github.com/google/uuid"
)

// DoesUserExist checks if userID entry exists in user table
func (p *PostgresDb) DoesUserExist(userID string) (exists bool, err error) {
	query :=
		`
			SELECT EXISTS(SELECT 1 FROM user_account WHERE user_id=$1)
		`

	err = p.dbConn.Get(&exists, query, userID)
	return
}

// CreateUser creates a new user and an accompanying organisation entry
func (p *PostgresDb) CreateUser(userID string, idType string) (orgID string, err error) {
	ctx := context.Background()
	tx, err := p.dbConn.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error begin trasacation to create user: %v\n", err.Error())
		return
	}

	insertUser :=
		`
			INSERT INTO user_account (user_id, id_type)
			VALUES ($1, $2)
		`
	_, err = tx.ExecContext(ctx, insertUser, userID, idType)
	if err != nil {
		log.Printf("Error insert user: %v\n", err.Error())
		tx.Rollback()
		return
	}

	insertOrg :=
		`
		INSERT INTO organisation (org_id, owner_id)
		VALUES ($1, $2)
		`
	orgID = uuid.NewString()
	_, err = tx.ExecContext(ctx, insertOrg, orgID, userID)
	if err != nil {
		log.Printf("Error insert organisation error: %v\n", err.Error())
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error commit transaction to create user: %v\n", err.Error())
		tx.Rollback()
		return
	}

	return
}
