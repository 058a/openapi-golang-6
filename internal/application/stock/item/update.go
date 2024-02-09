package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

func Update(req *UpdateRequestDto, r item.IRepository) error {
	// Precondition
	itemId, err := item.NewItemId(req.Id)
	if err != nil {
		return err
	}	

	itemName, err := item.NewItemName(req.Name)
	if err != nil {
		return err
	}

	// Main
	a, err := r.Get(itemId)
	if err != nil {
		return err
	}

	a.Name = itemName

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}