package item

type aggregate struct {
	id   		itemId
	Name 		itemName
	deleted bool
}

func NewAggregate(id itemId, name itemName) (*aggregate, error) {
	return &aggregate{
		id:   id,
		Name: name,
		deleted: false,
	}, nil
}

func (a aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *aggregate) Delete() {
	a.deleted = true
}