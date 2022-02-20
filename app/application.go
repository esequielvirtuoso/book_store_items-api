package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/esequielvirtuoso/book_store_items-api/internal/infrastructure/clients/elasticsearch"
	env "github.com/esequielvirtuoso/go_utils_lib/envs"
	"github.com/gorilla/mux"
)

const (
	envElasticURL     = "ES_URL"
	defaultElasticURL = "http://127.0.0.1:9200"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	env.CheckRequired(envElasticURL)

	elasticsearch.Init(getElasticSerachURL())

	mapUrls()

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8083",
		WriteTimeout: 500 * time.Microsecond,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Application Up!")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func getElasticSerachURL() string {
	return env.GetString(envElasticURL, defaultElasticURL)
}
