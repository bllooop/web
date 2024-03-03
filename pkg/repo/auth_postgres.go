package repo

import (
	"fmt"
	"web"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateShop(shop web.Shop) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (shopname, password) 
	VALUES ($1,$2) RETURNING id`, shopsTable)
	row := r.db.QueryRow(query, shop.ShopName, shop.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetShop(shopname, password string) (web.Shop, error) {
	var shop web.Shop
	query := fmt.Sprintf(`SELECT id FROM %s WHERE shopname=$1 AND password=$2`, shopsTable)
	err := r.db.Get(&shop, query, shopname, password)
	return shop, err
}
