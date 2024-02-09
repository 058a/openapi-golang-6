package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

func TestNewItemId(t *testing.T) {
	// When
	uuid := uuid.New()
	itemId, err := item.NewItemId(uuid)
	if err != nil {
		t.Fatal(err)
	}

	if itemId.UUID() != uuid {
		t.Errorf("itemId.UUID(%q) = %s; want %s", itemId.UUID(), uuid, uuid)
	}

	if itemId.String() != uuid.String() {
		t.Errorf("itemId.String(%q) = %s; want %s", itemId.String(), uuid.String(), uuid.String())
	}
}

func TestNewItemIdFailDueToEmpty(t *testing.T) {
	// When
	_, err := item.NewItemId(uuid.Nil)

	// Then
	if err != item.ErrItemIdEmpty {
		t.Errorf("expected %q, got %q", item.ErrItemIdEmpty, err)
	}
}