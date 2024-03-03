package service

import (
	"web"
	"web/pkg/repo"
)

type ProductListService struct {
	repo repo.ProductList
}

func NewProductListService(repo repo.ProductList) *ProductListService {
	return &ProductListService{repo: repo}
}

func (s *ProductListService) Create(shopId int, list web.ProductList) (int, error) {
	return s.repo.Create(shopId, list)
}

func (s *ProductListService) GetAllLists(shopId int) ([]web.ProductList, error) {
	return s.repo.GetAllLists(shopId)
}

func (s *ProductListService) GetById(shopId, listId int) (web.ProductList, error) {
	return s.repo.GetById(shopId, listId)
}

func (s *ProductListService) Delete(shopId, listId int) error {
	return s.repo.Delete(shopId, listId)
}
func (s *ProductListService) Update(shopId, listId int, input web.UpdateProductListInput) error {
	if err := input.Validation(); err != nil {
		return err
	}
	return s.repo.Update(shopId, listId, input)
}
