package location_test

import (
	"openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
	"testing"

	"github.com/google/uuid"
)

// テスト観点
// ・エラーが発生しないこと
// ・Update関数で作成したLocationをリポジトリで取得できること
// ・変更時に設定したLocationの名前が一致すること
func TestUpdate(t *testing.T) {
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
	beforeName := "TestName" + uuid.NewString()
	reqCreateDto := &location.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := location.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterName := "TestName" + uuid.NewString()
	reqUpdateDto := &location.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	err = location.Update(reqUpdateDto, repository)
	if err != nil {
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

	if a.Name.String() != afterName {
		t.Errorf("%T = %v, want %v", a.Name.String(), a.Name.String(), afterName)
	}
}
