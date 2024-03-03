package service

import (
	"web"
	"web/pkg/repo"
)

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
}
type Service struct {
	Authorization
	ProductList
	ProductItem
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ProductList:   NewProductListService(repos.ProductList),
	}
}
