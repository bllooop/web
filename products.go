package web

import "errors"

type ProductList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description"  db:"description"`
}

type ShopList struct {
	Id     int
	ShopId int
	ListId int
}

type ProductItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Expiration  string `json:"expiration" db:"expiration"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateProductListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateProductListInput) Validation() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update params have no values")
	}
	return nil
}

type UpdateProductItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	Expiration  *string `json:"expiration"`
}

func (i UpdateProductItemInput) Validation() error {
	if i.Title == nil && i.Description == nil && i.Price == nil && i.Expiration == nil {
		return errors.New("update params have no values")
	}
	return nil
}
