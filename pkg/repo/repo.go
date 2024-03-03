package repo

import (
	"web"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateShop(shop web.Shop) (int, error)
	GetShop(shopname, password string) (web.Shop, error)
}

type ProductList interface {
	Create(shopId int, list web.ProductList) (int, error)
	GetAllLists(shopId int) ([]web.ProductList, error)
	GetById(shopId, listId int) (web.ProductList, error)
	Delete(shopId, listId int) error
	Update(shopId, listId int, input web.UpdateProductListInput) error
}

type ProductItem interface {
}
type Repository struct {
	Authorization
	ProductList
	ProductItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProductList:   NewProductListPostgres(db),
	}
}
