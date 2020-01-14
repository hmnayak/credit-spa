package model

// Model stores a reference to a type that implements db interface
type Model struct {
	Db Db
}

// New returns a new instance of type Model
func New(db Db) *Model {
	return &Model{
		Db: db,
	}
}
