package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

func Delete(req *DeleteRequestDto, r item.IRepository) error {
	// Precondition
	itemId, err := item.NewItemId(req.Id)
	if err != nil {
		return err
	}	

	a, err := r.Get(itemId)
	if err != nil {
		return err
	}

	// Main
	a.Delete()

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}