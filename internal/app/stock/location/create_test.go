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
	reqDto := &app.CreateRequestDto{
		Name: "TestName" + uuid.NewString(),
	}

	// When
	resDto, err := app.Create(reqDto, repository, uuid.New())
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

func TestCreateFailInvalidName(t *testing.T) {
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
	reqDto := &app.CreateRequestDto{
		Name: "",
	}

	// When
	_, err = app.Create(reqDto, repository, uuid.New())

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestCreateFailInvalidId(t *testing.T) {
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
	reqDto := &app.CreateRequestDto{
		Name: "TestName" + uuid.NewString(),
	}

	// When
	_, err = app.Create(reqDto, repository, uuid.Nil)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestCreateFailSave(t *testing.T) {
	t.Parallel()

	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockIRepository(ctrl)
	repository.EXPECT().Save(gomock.Any()).Return(fmt.Errorf("fail save"))

	// Given
	reqDto := &app.CreateRequestDto{
		Name: "TestName" + uuid.NewString(),
	}

	// When
	_, err := app.Create(reqDto, repository, uuid.New())

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}
