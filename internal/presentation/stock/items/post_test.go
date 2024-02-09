package items_test

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"strings"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestPostCreated(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}

	name := uuid.NewString()

	// When
	postRes, err := rh.Post(
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: name,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	// Then
	if postRes.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBody, err := rch.AsCreated(postRes)
	if err != nil {
		t.Fatal(err)
	}

	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}
}

func TestPostBadRequest(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}
	
	zeroLenName := ""
	overLenName := strings.Repeat("a", 101)

	// When
	postResZeroLen, err := rh.Post(
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: zeroLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResZeroLen.Body.Close()

	postResOverLen, err := rh.Post(
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: overLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResOverLen.Body.Close()

	// Then
	if postResZeroLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResZeroLen.StatusCode)
	}	

	postResZeroLenBody, err := rch.AsBadRequest(postResZeroLen)
	if err != nil {
		t.Fatal(err)
	}
	if postResZeroLenBody.Message == "" {
		t.Errorf("expected empty, actual %s", postResZeroLenBody.Message)
	}
	
	if postResOverLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResOverLen.StatusCode)
	}	

	postResOverLenBody, err := rch.AsBadRequest(postResOverLen)
	if err != nil {
		t.Fatal(err)
	}
	if postResOverLenBody.Message== "" {
		t.Errorf("expected empty, actual %s", postResOverLenBody.Message)
	}
}