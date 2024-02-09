package location_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"openapi/internal/domain/stock/location"
	"openapi/internal/infrastructure/database"
	sut "openapi/internal/infrastructure/repository/sqlboiler/stock/location"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r, err := sut.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	currentDateTime := time.Now().UTC()

	// Given
	id, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := location.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	a := location.NewAggregate(id, name)

	// When
	before, err := r.Get(a.Id)
	if err == nil {
		t.Fatalf("expected error but returned nil, %+v", before)
	}

	if err = r.Save(a); err != nil {
		t.Fatal(err)
	}

	after, err := r.Get(a.Id)
	if err != nil {
		t.Fatalf("expected error but returned nil, %+v", err)
	}

	// Then
	if reflect.DeepEqual(after, before) {
		t.Errorf("%T %+v want %+v", after, after, before)
	}

	data, err := sqlboiler.FindStockLocation(context.Background(), db, a.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != id.String() {
		t.Errorf("%T %+v want %+v", data.ID, data.ID, id)
	}

	if data.Name != name.String() {
		t.Errorf("%T %+v want %+v", data.Name, data.Name, name)
	}

	if data.Deleted != false {
		t.Errorf("%T %+v want %+v", data.Deleted, data.Deleted, false)
	}

	if data.CreatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)
	}

	if data.UpdatedAt.Equal(data.CreatedAt) != true {
		t.Errorf("expected %s, got %s", data.CreatedAt, data.UpdatedAt)
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r, err := sut.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// Given
	id, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := location.NewName("before")
	if err != nil {
		t.Fatal(err)
	}

	before := location.NewAggregate(id, name)

	currentDateTime := time.Now().UTC()
	dataFormat := "2006-01-02 15:04:05.000000 +09:00"

	if err = r.Save(before); err != nil {
		t.Fatal(err)
	}

	beforeData, err := sqlboiler.FindStockLocation(context.Background(), db, before.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	after, err := r.Get(before.Id)
	if err != nil {
		t.Fatal(err)
	}

	changedName, err := location.NewName("after")
	if err != nil {
		t.Fatal(err)
	}

	after.Name = changedName
	after.Delete()

	if err = r.Save(after); err != nil {
		t.Fatal(err)
	}

	// Then
	afterData, err := sqlboiler.FindStockLocation(context.Background(), db, after.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	if afterData.ID != after.Id.String() {
		t.Errorf("%T %+v want %+v", afterData.ID, afterData.ID, after.Id.String())
	}

	if afterData.Name != after.Name.String() {
		t.Errorf("%T %+v want %+v", afterData.Name, afterData.Name, after.Name.String())
	}

	if afterData.Deleted != after.IsDeleted() {
		t.Errorf("%T %+v want %+v", afterData.Deleted, afterData.Deleted, after.IsDeleted())
	}

	if afterData.CreatedAt.Format(dataFormat) != beforeData.CreatedAt.Format(dataFormat) {
		t.Errorf("%T %+v want %+v", afterData.CreatedAt, afterData.CreatedAt, beforeData.CreatedAt.Format(dataFormat))
	}

	if afterData.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("%T %+v want greater than %+v ", afterData.UpdatedAt, afterData.UpdatedAt, currentDateTime)
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r, err := sut.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// Given
	id, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := location.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	a := location.NewAggregate(id, name)

	// When
	notFound, err := r.Find(a.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err = r.Save(a); err != nil {
		t.Fatal(err)
	}

	found, err := r.Find(a.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if notFound != false {
		t.Errorf("%T %+v want %+v", notFound, notFound, false)
	}

	if found != true {
		t.Errorf("%T %+v want %+v", found, found, true)
	}
}
