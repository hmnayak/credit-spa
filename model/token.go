package model

// User is a container of api user information used for login
type User struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

// AuthToken is a container of authentication token returned on successful login
type AuthToken struct {
	Token         string `json:"token"`
	UserName      string `json:"username"`
	Authorization string `db:"auth" json:"authorization"`
}
