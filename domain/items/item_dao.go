// Package items is the persistence layer of items domain.
package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/esequielvirtuoso/book_store_items-api/domain/queries"
	"github.com/esequielvirtuoso/book_store_items-api/internal/infrastructure/clients/elasticsearch"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() restErrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return restErrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() restErrors.RestErr {
	itemId := i.ID
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return restErrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return restErrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return restErrors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return restErrors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.ID = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, restErrors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, restErrors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, restErrors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.ID = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, restErrors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}

func (i *Item) Delete() restErrors.RestErr {

	_, err := elasticsearch.Client.Delete(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return restErrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return restErrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New("database error"))
	}

	return nil
}
