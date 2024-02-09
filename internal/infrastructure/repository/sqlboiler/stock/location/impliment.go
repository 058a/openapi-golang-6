package location

import (
	"context"
	"database/sql"
	"fmt"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"openapi/internal/domain/stock/location"
)

type Repository struct {
	location.IRepository
	db *sql.DB
}

func NewRepository(db *sql.DB) (*Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("NewRepository: db is nil")
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Save(a *location.Aggregate) error {
	data := &sqlboiler.StockLocation{
		ID:   a.Id.String(),
		Name: a.Name.String(),
		Deleted: a.IsDeleted(),
	}

	err := data.Upsert(
		context.Background(),
		r.db,
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

func (r *Repository) Get(id location.Id) (*location.Aggregate, error) {
	data, err := sqlboiler.FindStockLocation(context.Background(), r.db, id.UUID().String())
	if err != nil {
		return &location.Aggregate{}, err
	}

	name, err := location.NewName(data.Name)
	if err != nil {
		return &location.Aggregate{}, err
	}

	a := location.RestoreAggregate(id, name, data.Deleted)	

	return a, nil
}


func (r *Repository) Find(id location.Id) (bool, error) {
	found, err := sqlboiler.StockLocationExists(context.Background(), r.db, id.String())
	if err != nil {
		return false, err
	}
	
	return found, nil
}