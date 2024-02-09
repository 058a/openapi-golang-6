package location

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/location"
)

type CreateRequestDto struct {
	Name string
}

type CreateResponseDto struct {
	Id   uuid.UUID
	Name string
}

func Create(req *CreateRequestDto, r location.IRepository) (*CreateResponseDto, error) {
	// Precondition
	name, err := location.NewName(req.Name)
	if err != nil {
		return nil, err
	}

	// Main
	id, err := location.NewId(uuid.New())
	if err != nil {
		return nil, err
	}

	a := location.NewAggregate(id, name)

	if err := r.Save(a); err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id:   a.Id.UUID(),
		Name: a.Name.String(),
	}, nil
}
