package repo

import (
	"fmt"
	"strings"
	"web"

	"github.com/jmoiron/sqlx"
)

type ProductItemPostgres struct {
	db *sqlx.DB
}

func NewProductItemPostgres(db *sqlx.DB) *ProductItemPostgres {
	return &ProductItemPostgres{db: db}
}

// id title description price expiration

func (r *ProductItemPostgres) Create(listId int, product web.ProductItem) (int, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, price, expiration) VALUES ($1,$2,$3,$4) RETURNING ID", productItemTable)
	row := tr.QueryRow(createItemQuery, product.Title, product.Description, product.Price, product.Expiration)
	if err := row.Scan(&itemId); err != nil {
		tr.Rollback()
		return 0, err
	}
	createShopListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1,$2)", listsItemTable)
	_, err = tr.Exec(createShopListQuery, listId, itemId)
	if err != nil {
		tr.Rollback()
		return 0, err
	}
	return itemId, tr.Commit()

}

func (r *ProductItemPostgres) GetAll(shopId, listId int) ([]web.ProductItem, error) {
	var items []web.ProductItem
	query := fmt.Sprintf("SELECT pt.id, pt.title, pt.description, pt.price, pt.expiration FROM %s pt INNER JOIN %s li on li.item_id = pt.id INNER JOIN %s sl on sl.list_id = li.list_id WHERE li.list_id = $1 AND sl.shop_id = $2",
		productItemTable, listsItemTable, shopListTable)
	err := r.db.Select(&items, query, listId, shopId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ProductItemPostgres) GetById(shopId, itemId int) (web.ProductItem, error) {
	var product web.ProductItem

	query := fmt.Sprintf("SELECT pt.id, pt.title, pt.description, pt.price, pt.expiration FROM %s pt INNER JOIN %s li on li.item_id = pt.id INNER JOIN %s sl on sl.list_id = li.list_id WHERE pt.id = $1 AND sl.shop_id = $2",
		productItemTable, listsItemTable, shopListTable)
	err := r.db.Select(&product, query, itemId, shopId)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductItemPostgres) Delete(shopId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM  %s pi USING %s li, %s sl WHERE pt.id = li.item_id AND li.list_id = sl.list_id AND sl.list_id=$1 AND ti.ld=$2",
		productItemTable, listsItemTable, shopListTable)
	_, err := r.db.Exec(query, shopId, itemId)

	return err
}

func (r *ProductItemPostgres) Update(shopId, item_id int, input web.UpdateProductItemInput) error {
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
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	if input.Expiration != nil {
		setValues = append(setValues, fmt.Sprintf("expiration=$%d", argId))
		args = append(args, *input.Expiration)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s pt SET %s FROM %s li, %s sl WHERE ti.id = li.item_id AND li.list_id=sl.list_id AND sl.shop_id=$%d AND pt.id=$%d",
		productItemTable, setQuery, listsItemTable, shopListTable, argId, argId+1)
	args = append(args, shopId, item_id)
	_, err := r.db.Exec(query, args...)
	return err
}
