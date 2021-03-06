package store

import (
	"app/api"
	"app/api/pkg/mock"
)

// Factory is all of the stores in the application.
type Factory struct {
	User UserStore
	Note NoteStore
}

// NewFactory will return the factory for stores.
func NewFactory(m *mock.Mocker, db api.IDatabase, q api.IQuery) *Factory {
	cs := NewCore(m, db, q)

	return &Factory{
		User: NewUserStore(cs),
		Note: NewNoteStore(cs),
	}
}
