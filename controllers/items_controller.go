package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/esequielvirtuoso/book_store_items-api/domain/items"
	"github.com/esequielvirtuoso/book_store_items-api/domain/queries"
	"github.com/esequielvirtuoso/book_store_items-api/services"
	httputils "github.com/esequielvirtuoso/book_store_items-api/utils/http_utils"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
	"github.com/esequielvirtuoso/oauth_go_lib/oauth"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		httputils.RespondError(w, restErrors.NewUnauthorized("invalid request body"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondError(w, restErrors.NewBadRequestError("invalid request body"))
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		httputils.RespondError(w, restErrors.NewBadRequestError("invalid item json body"))
		return
	}

	itemRequest.SellerID = sellerId

	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		httputils.RespondError(w, createErr)
		return
	}

	httputils.RespondJson(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Get(itemId)
	if err != nil {
		httputils.RespondError(w, err)
		return
	}
	httputils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := restErrors.NewBadRequestError("invalid json body")
		httputils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := restErrors.NewBadRequestError("invalid json body")
		httputils.RespondError(w, apiErr)
		return
	}

	items, searchErr := services.ItemService.Search(query)
	if searchErr != nil {
		httputils.RespondError(w, searchErr)
		return
	}
	httputils.RespondJson(w, http.StatusOK, items)
}

func (cont *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		httputils.RespondError(w, restErrors.NewUnauthorized("invalid request body"))
		return
	}

	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Delete(itemId)
	if err != nil {
		httputils.RespondError(w, err)
		return
	}
	httputils.RespondJson(w, http.StatusOK, item)
}
