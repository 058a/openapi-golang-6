package main

import (
	"openapi/internal/infra/env"
	"testing"

	"net/http"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestMain(t *testing.T) {
	res, err := http.Get(env.GetServiceUrl())
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}
