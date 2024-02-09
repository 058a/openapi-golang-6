package location_test

import (
	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infrastructure/database"
	infra "openapi/internal/infrastructure/repository/sqlboiler/stock/location"
	"testing"

	"github.com/google/uuid"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repository, err := infra.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// Given
	reqCreateDto := &location.CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := location.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &location.DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	if err := location.Delete(reqDeleteDto, repository); err != nil {
		t.Fatal(err)
	}

	// Then
	id, err := domain.NewId(resCreateDto.Id)
	if err != nil {
		t.Fatal(err)
	}

	a, err := repository.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if !a.IsDeleted() {
		t.Errorf("%T = %v, want %v", a.IsDeleted(), a.IsDeleted(), false)
	}
}
