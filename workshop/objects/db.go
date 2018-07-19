package objects

import "errors"

// DB ...
type DB struct {
	data []string
}

// SaveAll ...
func (db DB) SaveAll(items []string) {
	for _, item := range items {
		db.Save(item)
	}
}

// Save ..
func (db DB) Save(item string) {
	db.data = append(db.data, item)
}

// Get ..
func (db DB) Get(item string) (string, error) {
	for _, v := range db.data {
		if v == item {
			return v, nil
		}
	}
	return "", errors.New("Not found")
}

// NewDB ..
func NewDB() *DB {
	db := new(DB)
	db.data = []string{}
	return db
}
