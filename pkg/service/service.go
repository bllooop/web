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
	}
}
