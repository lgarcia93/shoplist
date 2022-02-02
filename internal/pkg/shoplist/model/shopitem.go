package model

type ShopItem struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}
