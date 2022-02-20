// Package items is the persistence layer of items domain.
package items

import (
	"errors"

	"github.com/esequielvirtuoso/book_store_items-api/internal/infrastructure/clients/elasticsearch"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() restErrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return restErrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}
