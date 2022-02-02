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
	Create(item model.ShopItem) (int64, error)
	Update(item model.ShopItem) (int64, error)
	Delete(id int64) (int64, error)
	Get(id int64) (*model.ShopItem, error)
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

func (s shopListRepositoryImpl) Create(item model.ShopItem) (int64, error) {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return 0, fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	res, err := c.ExecContext(ctx, "insert into ShopItem(title, description, price) values(?, ?, ?)",
		item.Title,
		item.Description,
		item.Price,
	)

	if err != nil {
		fmt.Errorf("error insert data %v", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Errorf("error obtaining last insert ID %v", err)
	}

	return id, err

}

func (s shopListRepositoryImpl) Update(item model.ShopItem) (int64, error) {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return 0, fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	res, err := c.ExecContext(
		ctx,
		"Update ShopItem set title = ?, description = ?, price = ? where id = ? ",
		item.Title,
		item.Description,
		item.Price,
		item.ID,
	)

	if err != nil {
		return 0, fmt.Errorf("error updating the shopitem")
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf( "error obtaining affected rows")
	}

	return affectedRows, err
}

func (s shopListRepositoryImpl) Delete(id int64) (int64, error) {
	ctx := context.Background()
	c, err := s.connect(ctx)

	if err != nil {
		return 0, fmt.Errorf("error connecting to database")
	}

	defer c.Close()

	res, err := c.ExecContext(
		ctx,
		"Delete from ShopItem where id = ? ",
		id,
	)

	if err != nil {
		return 0, fmt.Errorf("error deleting the shopitem")
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf( "error obtaining affected rows")
	}

	return affectedRows, err
}

func (s shopListRepositoryImpl) Get(id int64) (*model.ShopItem, error) {
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
		return nil, fmt.Errorf("error reading record")
	}

	return &shopItem, nil
}

func (s shopListRepositoryImpl) GetAll() ([]*model.ShopItem, error) {
	ctx := context.Background()

	c, err := s.connect(ctx)

	if err != nil {
		return []*model.ShopItem{}, fmt.Errorf("error obtaining connection")
	}

	defer c.Close()

	rows, err := c.QueryContext(
		ctx,
		"select id, title, description, price from ShopItem",
	)

	var shopItems []*model.ShopItem

	for rows.Next() {
		var shopItem model.ShopItem

		if err := rows.Scan(
			&shopItem.ID,
			&shopItem.Title,
			&shopItem.Description,
			&shopItem.Price,
		); err != nil {
			return nil, fmt.Errorf("error reading records")
		}

		shopItems = append(shopItems, &shopItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving data")
	}

	return shopItems, nil
}

