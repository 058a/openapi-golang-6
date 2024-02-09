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
	reqCreateDto := &app.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := app.Create(reqCreateDto, repository, uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterName := "TestName" + uuid.NewString()
	reqUpdateDto := &app.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	err = app.Update(reqUpdateDto, repository)
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

func TestUpdateFailInvalidName(t *testing.T) {
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
	reqCreateDto := &app.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := app.Create(reqCreateDto, repository, uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterName := ""
	reqUpdateDto := &app.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}
	err = app.Update(reqUpdateDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestUpdateFailInvalidId(t *testing.T) {
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
	reqDto := &app.UpdateRequestDto{
		Id:   uuid.Nil,
		Name: "TestName" + uuid.NewString(),
	}

	// When
	err = app.Update(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestUpdateFailFind(t *testing.T) {
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
	reqDto := &app.UpdateRequestDto{
		Id:   uuid.New(),
		Name: "TestName" + uuid.NewString(),
	}

	// When
	err = app.Update(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}

func TestUpdateFailSave(t *testing.T) {
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
	reqDto := &app.UpdateRequestDto{
		Id:   id.UUID(),
		Name: "TestName" + uuid.NewString(),
	}

	// When
	err = app.Update(reqDto, repository)

	// Then
	if err == nil {
		t.Fatalf("error must not be nil")
	}
}
