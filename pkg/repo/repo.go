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
	}
}
