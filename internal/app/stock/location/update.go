package location

import (
	"openapi/internal/domain/stock/location"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string
}

func Update(req *UpdateRequestDto, r location.IRepository) error {
	// Precondition
	id, err := location.NewId(req.Id)
	if err != nil {
		return err
	}

	a, err := r.Get(id)
	if err != nil {
		return err
	}

	newName, err := location.NewName(req.Name)
	if err != nil {
		return err
	}

	// Main
	a.Name = newName

	if err = r.Save(a); err != nil {
		return err
	}

	return nil
}
