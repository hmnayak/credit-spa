package controller

import (
	"firebase.google.com/go/auth"
)

// Controller processes http requests.
// It contains a reference to the data store to perform CRUD operations.
type Controller struct {
	authSecret string
	authClient *auth.Client
}

// Init sets up a connection to database with configuration provided
func (c *Controller) Init(connStr string, authSecret string, authClient *auth.Client) error {
	c.authSecret = authSecret
	c.authClient = authClient

	return nil
}
