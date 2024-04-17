package service

import (
	"web"
	"web/pkg/repo"
)

// interface with service methods
type Authorization interface {
	CreateShop(shop web.Shop) (int, error)
	GenerateToken(shopname, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ProductList interface {
	Create(shopId int, list web.ProductList) (int, error)
	GetAllLists(shopId int) ([]web.ProductList, error)
	GetById(shopId, listId int) (web.ProductList, error)
	Delete(shopId, listId int) error
	Update(shopId, listid int, input web.UpdateProductListInput) error
}

type ProductItem interface {
	Create(shopId, listId int, product web.ProductItem) (int, error)
	GetAll(shopId, listId int) ([]web.ProductItem, error)
	GetById(shopId, itemId int) (web.ProductItem, error)
	Delete(shopId, itemId int) error
	Update(shopId, itemId int, input web.UpdateProductItemInput) error
}
type Service struct {
	Authorization
	ProductList
	ProductItem
}

// constructor
func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ProductList:   NewProductListService(repos.ProductList),
		ProductItem:   NewProductItemService(repos.ProductItem, repos.ProductList),
	}
}
