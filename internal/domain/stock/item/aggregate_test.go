package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}

	// When
	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.IsDeleted() != false {
		t.Errorf("expected %t, got %t", false, a.IsDeleted())
	}
}

func TestAggregateChangeName(t *testing.T) {
	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	beforeItemName, err := item.NewItemName("test1")
	if err != nil {
		t.Fatal(err)
	}

	a, err := item.NewAggregate(itemId, beforeItemName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterItemName, err := item.NewItemName("test2")
	if err != nil {
		t.Fatal(err)
	}

	a.Name = afterItemName

	// Then
	if a.Name.String() != "test2" {
		t.Errorf("a.Name.String(%q) = %s, want %s", a.Name.String(), a.Name.String(), "test2")
	}
}

func TestAggregateDelete(t *testing.T) {
	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}

	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.Delete()

	// Then
	if a.IsDeleted() != true {
		t.Errorf("a.IsDeleted() = %t, want %t", a.IsDeleted(), true)
	}
}