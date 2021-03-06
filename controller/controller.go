package controller

import (
	"firebase.google.com/go/auth"

	"github.com/hmnayak/credit/model"
)

// Controller processes http requests.
// It contains a reference to the data store to perform CRUD operations.
type Controller struct {
	model      *model.Model
	authSecret string
	authClient *auth.Client
}

// Init sets up a connection to database with configuration provided
func (c *Controller) Init(connStr string, authSecret string, authClient *auth.Client) error {
	c.authSecret = authSecret
	c.authClient = authClient

	return nil
}
