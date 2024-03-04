package service

import (
	"web"
	"web/pkg/repo"
)

type ProductItemService struct {
	repo     repo.ProductItem
	listRepo repo.ProductList
}

func NewProductItemService(repo repo.ProductItem, listRepo repo.ProductList) *ProductItemService {
	return &ProductItemService{repo: repo, listRepo: listRepo}
}

func (s *ProductItemService) Create(shopId, listId int, product web.ProductItem) (int, error) {
	_, err := s.listRepo.GetById(shopId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, product)
}

func (s *ProductItemService) GetAll(shopId, listId int) ([]web.ProductItem, error) {
	return s.repo.GetAll(shopId, listId)
}

func (s *ProductItemService) GetById(shopId, itemId int) (web.ProductItem, error) {
	return s.repo.GetById(shopId, itemId)
}

func (s *ProductItemService) Delete(shopId, itemId int) error {
	return s.repo.Delete(shopId, itemId)
}

func (s *ProductItemService) Update(shopId, item_id int, input web.UpdateProductItemInput) error {
	if err := input.Validation(); err != nil {
		return err
	}
	return s.repo.Update(shopId, item_id, input)
}
