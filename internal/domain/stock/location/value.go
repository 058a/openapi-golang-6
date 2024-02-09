package location

import "fmt"

type Name struct {
	string
}

func NewName(v string) (Name, error) {
	if v == "" {
		return Name{}, fmt.Errorf("NewName: invalid name %+v", v)
	}
	return Name{v}, nil
}

func (v Name) String() string {
	return v.string
}