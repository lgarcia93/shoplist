package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/model"
)

type repository interface {
	connect(ctx context.Context) (*sql.Conn, error)
}

type ShopListRepository interface {
	Create(item model.ShopItem) error
	Update(item model.ShopItem) error
	Delete(id int64) error
	GetByID(id int64) (*model.ShopItem, error)
	GetAll() ([]*model.ShopItem, error)
}

type shopListRepositoryImpl struct {
	db *sql.DB
}

func NewShopListRepository(db *sql.DB) ShopListRepository {
	return shopListRepositoryImpl{
		db: db,
	}
}

func (s shopListRepositoryImpl) connect(ctx context.Context) (*sql.Conn, error) {
	return s.db.Conn(ctx)
}

func (s shopListRepositoryImpl) Create(item model.ShopItem) error {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	prepared, err := c.PrepareContext(ctx, "Insert into ShopItem(title, description, price) values(?, ?, ?)")

	_, err = prepared.ExecContext(
		ctx,
		item.Title,
		item.Description,
		item.Price,
	)

	return err
}

func (s shopListRepositoryImpl) Update(item model.ShopItem) error {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	stmt, err := c.PrepareContext(ctx, "Update ShopItem set title = ?, description = ?, price = ? where id = ?")

	res, err := stmt.ExecContext(
		ctx,
		item.Title,
		item.Description,
		item.Price,
		item.ID,
	)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	if err != nil {
		return fmt.Errorf("error obtaining affected rows")
	}

	return err
}

func (s shopListRepositoryImpl) Delete(id int64) error {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	prepared, err := c.PrepareContext(ctx, "Delete from ShopItem where id = ?")

	if err != nil {
		return fmt.Errorf("error preparing statement")
	}

	_, err = prepared.ExecContext(ctx, id)

	if err != nil {
		return fmt.Errorf("error deleting the shopitem")
	}

	return err
}

func (s shopListRepositoryImpl) GetByID(id int64) (*model.ShopItem, error) {
	ctx := context.Background()

	c, err := s.connect(ctx)

	if err != nil {
		return nil, fmt.Errorf("error obtaining connection")
	}

	defer c.Close()

	row := c.QueryRowContext(
		ctx,
		"Select id, title, description, price from ShopItem where id = ?",
		id,
	)

	var shopItem model.ShopItem

	err = row.Scan(
		&shopItem.ID,
		&shopItem.Title,
		&shopItem.Description,
		&shopItem.Price,
	)

	if err != nil {
		return nil, err
	}

	return &shopItem, nil
}

func (s shopListRepositoryImpl) GetAll() ([]*model.ShopItem, error) {
	ctx := context.Background()

	c, err := s.connect(ctx)

	if err != nil {
		return make([]*model.ShopItem, 0), fmt.Errorf("error obtaining connection")
	}

	defer c.Close()

	rows, err := c.QueryContext(
		ctx,
		"Select id, title, description, price from ShopItem",
	)

	if err != nil {
		return nil, err
	}

	shopItems := make([]*model.ShopItem, 0)

	for rows.Next() {
		var shopItem model.ShopItem

		if err := rows.Scan(
			&shopItem.ID,
			&shopItem.Title,
			&shopItem.Description,
			&shopItem.Price,
		); err != nil {
			return nil, err
		}

		shopItems = append(shopItems, &shopItem)
	}

	if err := rows.Err(); err != nil {
		return shopItems, err
	}

	return shopItems, nil
}
