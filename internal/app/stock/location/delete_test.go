package location_test

import (
	"fmt"
	app "openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	mock "openapi/internal/infra/mock/domain/stock/location"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
	"testing"

	"github.com/golang/mock/gomock"
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
	reqCreateDto := &app.CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := app.Create(reqCreateDto, repository, uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &app.DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	if err := app.Delete(reqDeleteDto, repository); err != nil {
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

func TestDeleteFailInvalidId(t *testing.T) {
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
	reqDto := &app.DeleteRequestDto{
		Id: uuid.Nil,
	}

	// When
	err = app.Delete(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestDeleteFailFind(t *testing.T) {
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
	reqDto := &app.DeleteRequestDto{
		Id: uuid.New(),
	}

	// When
	err = app.Delete(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestDeleteFailSave(t *testing.T) {
	t.Parallel()

	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockIRepository(ctrl)

	repository.EXPECT().Save(gomock.Any()).Return(fmt.Errorf("fail save"))

	name, err := domain.NewName("TestName" + uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	a := domain.NewAggregate(id, name)
	repository.EXPECT().Get(gomock.Any()).Return(a, nil)

	// Given
	reqDto := &app.DeleteRequestDto{
		Id: id.UUID(),
	}

	// When
	err = app.Delete(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}
