package repo

import (
	"fmt"
	"strings"
	"web"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductListPostgres struct {
	db *sqlx.DB
}

func NewProductListPostgres(db *sqlx.DB) *ProductListPostgres {
	return &ProductListPostgres{db: db}
}

func (r *ProductListPostgres) Create(shopId int, list web.ProductList) (int, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1,$2) RETURNING ID", productListTable)
	row := tr.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tr.Rollback()
		return 0, err
	}
	createShopListQuery := fmt.Sprintf("INSERT INTO %s (shop_id, list_id) VALUES ($1,$2)", shopListTable)
	_, err = tr.Exec(createShopListQuery, shopId, id)
	if err != nil {
		tr.Rollback()
		return 0, err
	}
	return id, tr.Commit()

}

func (r *ProductListPostgres) GetAllLists(shopId int) ([]web.ProductList, error) {
	var lists []web.ProductList
	query := fmt.Sprintf("SELECT pl.id, pl.title, pl.description FROM %s pl INNER JOIN %s sl on pl.id = sl.list_id WHERE sl.shop_id = $1",
		productListTable, shopListTable)
	err := r.db.Select(&lists, query, shopId)

	return lists, err
}

func (r *ProductListPostgres) GetById(shopId, listId int) (web.ProductList, error) {
	var list web.ProductList

	query := fmt.Sprintf("SELECT pl.id, pl.title, pl.description FROM %s pl INNER JOIN %s sl on pl.id = sl.list_id WHERE sl.shop_id = $1 AND sl.list_id=$2",
		productListTable, shopListTable)
	err := r.db.Select(&list, query, shopId, listId)

	return list, err
}

func (r *ProductListPostgres) Delete(shopId, listId int) error {
	query := fmt.Sprintf("DELETE FROM  %s pl USING %s sl WHERE  pl.id = sl.list_id AND sl.shop_id = $1 AND sl.list_id=$2",
		productListTable, shopListTable)
	_, err := r.db.Exec(query, shopId, listId)

	return err
}

func (r *ProductListPostgres) Update(shopId, listId int, input web.UpdateProductListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s pl SET %s FROM %s sl WHERE pl.id = sl.list_id AND sl.list_id=$%d AND sl.shop_id=$%d",
		productListTable, setQuery, shopListTable, argId, argId+1)
	args = append(args, listId, shopId)
	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}
