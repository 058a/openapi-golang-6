package location

import (
	"fmt"

	"github.com/google/uuid"
)

type Id struct {
	value uuid.UUID
}

func NewId(v uuid.UUID) (Id, error) {
	if v == uuid.Nil {
		return Id{}, fmt.Errorf("invalid id because empty")
	}
	return Id{v}, nil
}

func (v Id) UUID() uuid.UUID {
	return v.value
}

func (v Id) String() string {
	return v.value.String()
}