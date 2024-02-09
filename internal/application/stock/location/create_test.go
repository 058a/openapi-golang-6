package location_test

import (
	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infrastructure/database"
	infra "openapi/internal/infrastructure/repository/sqlboiler/stock/location"
	"testing"

	"github.com/google/uuid"
)

// テスト観点
// ・エラーが発生しないこと
// ・Create関数で作成したLocationをリポジトリで取得できること
// ・作成時に設定したLocationの名前が一致すること
func TestCreate(t *testing.T) {
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
	reqDto := &location.CreateRequestDto{
		Name: "TestName" + uuid.NewString(),
	}

	// When
	resDto, err := location.Create(reqDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	id, err := domain.NewId(resDto.Id)
	if err != nil {
		t.Fatal(err)
	}

	a, err := repository.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if a.Name.String() != reqDto.Name {
		t.Errorf("%T = %v, want %v", a.Name.String(), a.Name.String(), reqDto.Name)
	}
}
