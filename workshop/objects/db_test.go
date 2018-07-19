package objects

import (
	"testing"
)

// Save ...
func TestSaveDb(t *testing.T) {
	items := []string{"test1", "test2", "test3"}

	db := NewDB()

	db.SaveAll(items)
	value, err := db.Get(items[0])

	if err != nil {
		t.Errorf("Failed to get %q from db", items[0])
	}

	if value != items[0] {
		t.Errorf("expected %q got %q", items[0], value)
	}
}
