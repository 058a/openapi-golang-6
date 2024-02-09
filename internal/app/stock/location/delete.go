package location

import (
	"openapi/internal/domain/stock/location"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

func Delete(req *DeleteRequestDto, r location.IRepository) error {
	// Precondition
	id, err := location.NewId(req.Id)
	if err != nil {
		return err
	}

	a, err := r.Get(id)
	if err != nil {
		return err
	}

	// Main
	a.Delete()

	if err = r.Save(a); err != nil {
		return err
	}

	return nil
}
