package item_test

import (
	"context"
	"testing"
	"time"

	"openapi/internal/domain/stock/item"

	"openapi/internal/infrastructure/database"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	} 
	itemName, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}
	currentDateTime := time.Now().UTC()

	// When
	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	data, err := sqlboiler.FindStockItem(context.Background(), db, itemId.UUID().String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != itemId.String() {
		t.Errorf("data.ID(%q) = %s; want %s", data.ID, data.ID, itemId.String())
	}
	
	if data.Name != itemName.String() {
		t.Errorf("data.Name(%q) = %s; want %s", data.Name, data.Name, itemName.String())
	}

	if data.Deleted != false {
		t.Errorf("data.Deleted(%t) = %t; want %t", data.Deleted, data.Deleted, false)
	}

	if data.CreatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)		
	}

	if data.UpdatedAt.Equal(data.CreatedAt) != true {
		t.Errorf("expected %s, got %s", data.CreatedAt, data.UpdatedAt)
	}
}

func TestUpdate(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	beforeItemName, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}

	a, err := item.NewAggregate(itemId, beforeItemName)
	if err != nil {
		t.Fatal(err)
	}

	currentDateTime := time.Now().UTC()
	dataFormat := "2006-01-02 15:04:05.000000 +09:00"

	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	beforeData, err := sqlboiler.FindStockItem(context.Background(), db, itemId.String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterItemName, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}

	a.Name = afterItemName
	a.Delete()

	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	afterData, err := sqlboiler.FindStockItem(context.Background(), db, itemId.String())
	if err != nil {
		t.Fatal(err)
	}

	if afterData.ID != beforeData.ID{
		t.Errorf("afterData.ID(%q) = %s; want %s", afterData.ID, afterData.ID, beforeData.ID)
	}

	if afterData.Name != afterItemName.String() {
		t.Errorf("afterData.Name(%q) = %s; want %s", afterData.Name, afterData.Name, afterItemName.String())
	}

	if afterData.Deleted != a.IsDeleted() {
		t.Errorf("afterData.Deleted(%t) = %t; want %t", afterData.Deleted, afterData.Deleted, a.IsDeleted())
	}

	if afterData.CreatedAt.Format(dataFormat) != beforeData.CreatedAt.Format(dataFormat) {
		t.Errorf("afterData.CreatedAt(%s) = %s; want %s", afterData.CreatedAt, afterData.CreatedAt, beforeData.CreatedAt)
	}

	if afterData.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("afterData.UpdatedAt(%s) = %s; want %s", afterData.UpdatedAt, afterData.UpdatedAt, currentDateTime)
	}
}

func TestFind(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}

	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	beforeFound, err := r.Find(itemId)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	afterFound, err := r.Find(itemId)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if beforeFound != false {
		t.Errorf("expected %t, got %t", false, beforeFound)
	}
	if afterFound != true {
		t.Errorf("expected %t, got %t", true, afterFound)
	}
}