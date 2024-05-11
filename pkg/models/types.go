package models

// Entity - Base interface for all the entities which is going to save in database
type Entity interface {
	GetId() int
	SetId(id int) Entity
}
