package item

import (
	"fmt"

	"github.com/google/uuid"
)

type itemId struct {
	value uuid.UUID
}

var ErrItemIdEmpty = fmt.Errorf("invalid itemId because empty")

func NewItemId(v uuid.UUID) (itemId, error) {
	if v == uuid.Nil {
		return itemId{}, ErrItemIdEmpty
	}
	
	return itemId{v}, nil
}

func (e itemId) UUID() uuid.UUID {
	return e.value
}

func (e itemId) String() string {
	return e.value.String()
}