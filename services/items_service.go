package services

import (
	"net/http"

	"github.com/esequielvirtuoso/book_store_items-api/domain/items"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

var (
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, restErrors.RestErr)
	Get(string) (*items.Item, restErrors.RestErr)
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
	return nil, restErrors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
