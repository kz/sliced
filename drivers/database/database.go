package database

type Database interface {
	AddItem(map[string]interface{})
}
