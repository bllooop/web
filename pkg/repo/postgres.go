package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	listsItemTable   = "listsItem"
	shopListTable    = "shoplist"
	productListTable = "productlist"
	shopsTable       = "shops"
	productItemTable = "productItem"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
