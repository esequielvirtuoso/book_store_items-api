package app

import (
	"net/http"

	"github.com/esequielvirtuoso/book_store_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodDelete)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
