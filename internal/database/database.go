package database

import "fmt"

type Store struct {
	db_name string
}

func NewDB(name string) *Store {
	return &Store{
		db_name: name,
	}
}

func (s *Store) DoDbStuff() {
	fmt.Printf("I'm getting something from %s, hold your horses.", s.db_name)
}
