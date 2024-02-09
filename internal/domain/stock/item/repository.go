package item

import (
	"context"
	"database/sql"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(a *aggregate) error
	Get(id itemId) (*aggregate, error)
	Find(id itemId) (bool, error)
}

type Repository struct {
	IRepository
	Db *sql.DB
}

func (r *Repository) Save(a *aggregate) error {
	data := &sqlboiler.StockItem{
		ID:   a.id.UUID().String(),
		Name: a.Name.value,
		Deleted: a.deleted,
	}

	err := data.Upsert(
		context.Background(),
		r.Db,
		true,
		[]string{"id"},
		boil.Whitelist("name","deleted"),
		boil.Infer(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(id itemId) (*aggregate, error) {
	data, err := sqlboiler.FindStockItem(context.Background(), r.Db, id.String())
	if err != nil {
		return &aggregate{}, err
	}

	itemName, err := NewItemName(data.Name)
	if err != nil {
		return &aggregate{}, err
	}

	a := &aggregate{
		id:   id,
		Name: itemName,
		deleted: data.Deleted,
	}
	
	return a, nil
}


func (r *Repository) Find(id itemId) (bool, error) {
	found, err := sqlboiler.StockItemExists(context.Background(), r.Db, id.String())
	if err != nil {
		return false, err
	}
	
	return found, nil
}