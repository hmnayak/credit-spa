package model

// User is a container of api user information used for login
type User struct {
	UserID string `db:"user_id" json:"userid"`
}

// AuthToken is a container of authentication token returned on successful login
type AuthToken struct {
	Token         string `json:"token"`
	UserName      string `json:"username"`
	Authorization string `db:"auth" json:"authorization"`
}
