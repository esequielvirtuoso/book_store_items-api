package services

import (
	"github.com/esequielvirtuoso/book_store_items-api/domain/items"
	"github.com/esequielvirtuoso/book_store_items-api/domain/queries"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

var (
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, restErrors.RestErr)
	Get(string) (*items.Item, restErrors.RestErr)
	Search(queries.EsQuery) ([]items.Item, restErrors.RestErr)
	Delete(string) (*items.Item, restErrors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, restErrors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil

}

func (s *itemsService) Get(ID string) (*items.Item, restErrors.RestErr) {
	item := items.Item{ID: ID}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, restErrors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemsService) Delete(ID string) (*items.Item, restErrors.RestErr) {
	item := items.Item{ID: ID}

	if err := item.Delete(); err != nil {
		return nil, err
	}
	return nil, nil
}
