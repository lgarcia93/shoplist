package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/model"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/repository"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Error opening mock connection")
	}

	defer db.Close()

	repo := repository.NewShopListRepository(db)

	query := "Select id, title, description, price from ShopItem where id = ?"

	mockShopItem := model.ShopItem{
		ID:          1,
		Title:       "NewItem",
		Description: "NewItemForTEst",
		Price:       25.0,
	}

	rows := mock.NewRows([]string{"id", "title", "description", "price"}).
		AddRow(
			mockShopItem.ID,
			mockShopItem.Title,
			mockShopItem.Description,
			mockShopItem.Price,
		)

	mock.ExpectQuery(query).WithArgs(mockShopItem.ID).WillReturnRows(rows)

	shopitem, err := repo.GetByID(1)

	assert.NotNil(t, shopitem)
	assert.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Error opening mock connection")
	}

	defer db.Close()

	repo := repository.NewShopListRepository(db)

	query := regexp.QuoteMeta("Select id, title, description, price from ShopItem")

	mockShopItems := []model.ShopItem{
		{
			ID:          1,
			Title:       "Item1",
			Description: "Item1 Description",
			Price:       32.5,
		},
		{
			ID:          2,
			Title:       "Item2",
			Description: "Item2 Description",
			Price:       14.5,
		},
		{
			ID:          3,
			Title:       "Item3",
			Description: "Item3 Description",
			Price:       5.99,
		},
	}

	rows := mock.NewRows([]string{"id", "title", "description", "price"})

	for _, v := range mockShopItems {
		rows.AddRow(
			v.ID,
			v.Title,
			v.Description,
			v.Price,
		)
	}

	mock.ExpectQuery(query).WillReturnRows(rows)

	shopItems, err := repo.GetAll()

	assert.NotEmpty(t, shopItems)
	assert.NoError(t, err)
	assert.Len(t, shopItems, len(mockShopItems))
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Error opening mock connection")
	}

	defer db.Close()

	query := regexp.QuoteMeta("Update ShopItem set title = ?, description = ?, price = ? where id = ?")

	shopItem := model.ShopItem{
		ID:          1,
		Title:       "Item 1",
		Description: "Item Description 1",
		Price:       20.0,
	}

	prepared := mock.ExpectPrepare(query)

	prepared.ExpectExec().WithArgs(
		shopItem.Title,
		shopItem.Description,
		shopItem.Price,
		shopItem.ID,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewShopListRepository(db)

	err = repo.Update(shopItem)

	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error opening connection")
	}

	defer db.Close()

	query := "Insert into ShopItem(title, description, price) values(?, ?, ?)"

	prepared := mock.ExpectPrepare(query)

	shopItem := model.ShopItem{
		Title:       "Created Item in Tests",
		Description: "Mocked description",
		Price:       34.4,
	}

	prepared.ExpectExec().WithArgs(
		shopItem.Title,
		shopItem.Description,
		shopItem.Price,
	).WillReturnResult(sqlmock.NewResult(0, 0))

	repo := repository.NewShopListRepository(db)

	err = repo.Create(shopItem)

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error opening connection")
	}

	defer db.Close()

	query := "Delete from ShopItem where id = ?"

	prepared := mock.ExpectPrepare(query)

	var shopListId int64 = 1

	prepared.ExpectExec().
		WithArgs(shopListId).
		WillReturnResult(
			sqlmock.NewResult(
				0,
				1,
			),
		)

	repo := repository.NewShopListRepository(db)

	repo.Delete(shopListId)

	assert.NoError(t, err)
}
