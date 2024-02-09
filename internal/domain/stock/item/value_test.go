package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"
)

func TestNewItemName(t *testing.T) {
	// Given
	name := "test"

	// When
	a, err := item.NewItemName(name)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.String() != name {
		t.Errorf("expected %s, got %s", name, a.String())
	}
}

func TestNewItemNameFailDueToEmpty(t *testing.T) {
	// Given
	name := ""

	// When
	_, err := item.NewItemName(name)

	// Then
	if err != item.ErrItemNameEmpty {
		t.Fatalf("expected %q, got %q", item.ErrItemNameEmpty, err)
	}
}