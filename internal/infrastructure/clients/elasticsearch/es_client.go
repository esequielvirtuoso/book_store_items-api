package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/esequielvirtuoso/go_utils_lib/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init(esURL string) {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)

	// Create the index if it does not exists
	// I.e:
	// PUT
	// {
	// 	"settings": {
	// 		"index": {
	// 			"number_of_shards": 4,
	// 			"number_of_replicas": 2
	// 		}
	// 	}
	// }
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, document interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index("items").BodyJson(document).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}
