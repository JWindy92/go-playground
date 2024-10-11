package models

import "fmt"

type DatabaseInterface interface {
	DoDbStuff()
}

type ModelsImplementation struct {
	database DatabaseInterface
}

func NewModelsImplementation(db DatabaseInterface) *ModelsImplementation {
	return &ModelsImplementation{
		database: db,
	}
}

func (m *ModelsImplementation) GetSomethingFromDb() {
	fmt.Println("Models speaking. Oh, it's you again. The usual? Okay I'll ask my database...")
	m.database.DoDbStuff()
}
