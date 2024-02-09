package item

import "fmt"

type itemName struct {
	value string
}

var ErrItemNameEmpty = fmt.Errorf("invalid itemName because empty")

func NewItemName(v string) (itemName, error) {
	if v == "" {
		return itemName{}, ErrItemNameEmpty
	}
	return itemName{v}, nil
}

func (v itemName) String() string {
	return v.value
}