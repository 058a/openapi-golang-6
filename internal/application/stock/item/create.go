package item

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/item"
)

type CreateRequestDto struct {
	Name string	
}

type CreateResponseDto struct {
	Id uuid.UUID
	Name string	
}

func Create(req *CreateRequestDto, r item.IRepository) (*CreateResponseDto, error) {
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		return nil, err
	}
	itemName, err := item.NewItemName(req.Name)
	if err != nil {
		return nil, err
	}
	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		return nil, err
	}

	if err := r.Save(a); err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: itemId.UUID(),
		Name: itemName.String(),
	}, nil
}