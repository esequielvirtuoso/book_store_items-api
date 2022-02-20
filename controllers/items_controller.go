package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/esequielvirtuoso/book_store_items-api/domain/items"
	"github.com/esequielvirtuoso/book_store_items-api/services"
	httputils "github.com/esequielvirtuoso/book_store_items-api/utils/http_utils"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
	"github.com/esequielvirtuoso/oauth_go_lib/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// httputils.RespondError(w, err)
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

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
