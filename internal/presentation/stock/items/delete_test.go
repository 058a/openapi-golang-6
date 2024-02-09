package items_test

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestDeleteOk(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}

	// Given
	postRes, err := rh.Post(
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: uuid.NewString(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatalf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBody, err := rch.AsCreated(postRes)
	if err != nil {
		t.Fatal(err)
	}

	// When
	deleteRes, err := rh.Delete(postResBody.Id)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	// Then
	if deleteRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, deleteRes.StatusCode)
	}
}

func TestDeleteNotFound(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}

	deleteRes, err := rh.Delete(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	// Then
	if deleteRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, deleteRes.StatusCode)
	}
}